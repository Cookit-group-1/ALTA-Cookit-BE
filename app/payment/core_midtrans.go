package payment

import (
	"alta-cookit-be/app/config"
	"alta-cookit-be/features/transaction_details"
	"errors"
	"log"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentGateway interface {
	ChargeTransaction(orderID string, bank string, transactionDetailEntities []transaction_details.TransactionDetailEntity) (string, error)
}

type midtransCore struct {
	core coreapi.Client
}

func NewCoreMidtrans() PaymentGateway {
	c := coreapi.Client{}
	c.New(config.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	return midtransCore{core: c}
}

// 'None', 'COD', 'SeaBank', 'BCA', 'BNI', 'Mandiri', 'QRIS', 'Gopay'
func (c midtransCore) ChargeTransaction(orderID string, bank string, transactionDetailEntities []transaction_details.TransactionDetailEntity) (string, error) {
	grossAmt, midTransItemDetails := int64(0), []midtrans.ItemDetails{}
	for _, entity := range transactionDetailEntities {
		grossAmt += int64(entity.Price)
		midTransItemDetails = append(midTransItemDetails, midtrans.ItemDetails{
			ID:    "Ingredient-" + strconv.Itoa(int(entity.ID)),
			Qty:   int32(entity.Quantity),
			Price: int64(entity.Price)/int64(entity.Quantity),
			Name:  entity.IngredientName,
		})
	}

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(grossAmt),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(bank),
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 1,
			Unit:           "day",
		},
		Items: &midTransItemDetails,
	}

	response, errMidtrans := c.core.ChargeTransaction(chargeReq)
	if errMidtrans != nil {
		log.Println(errMidtrans)
		return "", errors.New("charge transaction failed due to internal server error")
	}

	return response.VaNumbers[0].VANumber, nil
}
