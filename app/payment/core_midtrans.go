package payment

import (
	"alta-cookit-be/app/config"
	"alta-cookit-be/features/transactions"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentGateway interface {
	ChargeTransaction(transactions.TransactionEntity) (string, error)
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
func (c midtransCore) ChargeTransaction(transactionEntity transactions.TransactionEntity) (string, error) {
	grossAmt, midTransItemDetails := int64(0), []midtrans.ItemDetails{}
	for _, entity := range transactionEntity.TransactionDetailEntities {
		grossAmt += int64(entity.Price)
		midTransItemDetails = append(midTransItemDetails, midtrans.ItemDetails{
			ID:    "Ingredient-" + strconv.Itoa(int(entity.ID)),
			Qty:   int32(entity.Quantity),
			Price: int64(entity.Price) / int64(entity.Quantity),
			Name:  entity.IngredientName,
		})
	}
	grossAmt += int64(transactionEntity.ShippingFee)
	midTransItemDetails = append(midTransItemDetails, midtrans.ItemDetails{
		ID:    "Shipping Fee-" + strconv.Itoa(int(transactionEntity.ID)),
		Qty:   1,
		Price: int64(transactionEntity.ShippingFee),
		Name:  transactionEntity.ShippingMethod,
	})

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionEntity.OrderID,
			GrossAmt: int64(grossAmt),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(strings.ToLower(transactionEntity.PaymentMethod)),
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
