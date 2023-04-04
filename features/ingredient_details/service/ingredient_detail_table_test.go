package service

import (
	"alta-cookit-be/features/ingredient_details"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		ingredientDetailEntity ingredient_details.IngredientDetailEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		result        *ingredient_details.IngredientDetailEntity
		errResult     error
	}
}

func InsertIngredientDetailTestTable() []TestTable {
	tname := "test insert ingredient detail "
	return []TestTable{
		{
			Name: tname + "expect failed - empty ingredient detail name",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
			Name: tname + "expect failed - empty ingredient detail quantity",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
			Name: tname + "expect success - filled ingredient detail name and unit",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
				errResult: nil,
			},
		},
	}
}

func UpdateIngredientDetailByIdTestTable() []TestTable {
	tname := "test update ingredient detail by id "
	return []TestTable{
		{
			Name: tname + "expect failed - empty ingredient detail name",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
			Name: tname + "expect failed - empty ingredient detail quantity",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
			Name: tname + "expect success - filled ingredient detail name and unit",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
				errResult: nil,
			},
		},
	}
}

func DeleteIngredientDetailByIdTestTable() []TestTable {
	tname := "test delete ingredinet detail by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
				ingredientDetailEntity ingredient_details.IngredientDetailEntity
			}{
				ingredientDetailEntity: ingredient_details.IngredientDetailEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Kentang",
					Quantity: 1,
					Unit:     "pcs",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredient_details.IngredientDetailEntity
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
