package service

import (
	"alta-cookit-be/features/carts"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		cartEntity carts.CartEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		result        *carts.CartEntity
		results       *[]carts.CartEntity
		errResult     error
	}
}

func SelectCartsByUserIdTable() []TestTable {
	tname := "test select carts by user id "
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result:        nil,
				errResult:     nil,
			},
		},
	}
}

func InsertCartTestTable() []TestTable {
	tname := "test insert cart "
	return []TestTable{
		{
			Name: tname + "expect failed - empty quantity",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: true,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.VALIDATION_InvalidInput),
			},
		},
		{
			Name: tname + "expect success - filled quantity",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
				errResult: nil,
			},
		},
	}
}

func UpdateCartByIdTestTable() []TestTable {
	tname := "test update cart by id "
	return []TestTable{
		{
			Name: tname + "expect failed - empty quantity",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: true,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.VALIDATION_InvalidInput),
			},
		},
		{
			Name: tname + "expect success - filled quantity",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
				errResult: nil,
			},
		},
	}
}

func DeleteCartByIdTestTable() []TestTable {
	tname := "test delete cart by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				result:        nil,
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				cartEntity carts.CartEntity
			}{
				cartEntity: carts.CartEntity{
					ID:           1,
					UserID:       1,
					IngredientID: 1,
					Quantity:     1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *carts.CartEntity
				results       *[]carts.CartEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result:        nil,
				errResult:     nil,
			},
		},
	}
}
