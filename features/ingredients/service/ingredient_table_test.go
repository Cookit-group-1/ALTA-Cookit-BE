package service

import (
	"alta-cookit-be/features/ingredients"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		ingredientEntity ingredients.IngredientEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		result        *ingredients.IngredientEntity
		errResult     error
	}
}

func InsertIngredientTestTable() []TestTable {
	tname := "test insert ingredient "
	return []TestTable{
		{
			Name: tname + "expect failed - empty ingredient name",
			Input: struct {
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
			Name: tname + "expect success - filled ingredient name",
			Input: struct {
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &ingredients.IngredientEntity{
					ID:       1,
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
				errResult: nil,
			},
		},
	}
}

func UpdateIngredientByIdTestTable() []TestTable {
	tname := "test update ingredient by id "
	return []TestTable{
		{
			Name: tname + "expect failed - empty ingredient name",
			Input: struct {
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
			Name: tname + "expect success - filled step name",
			Input: struct {
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &ingredients.IngredientEntity{
					ID:       1,
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
				errResult: nil,
			},
		},
	}
}

func DeleteIngredientByIdTestTable() []TestTable {
	tname := "test delete ingredinet by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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

func DeleteIngredientByRecipeIdTestTable() []TestTable {
	tname := "test delete ingredient by recipe id"
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
				ingredientEntity ingredients.IngredientEntity
			}{
				ingredientEntity: ingredients.IngredientEntity{
					UserID:   1,
					UserRole: "User",
					RecipeID: 1,
					Name:     "Original",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *ingredients.IngredientEntity
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
