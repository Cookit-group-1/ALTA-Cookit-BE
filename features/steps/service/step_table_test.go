package service

import (
	"alta-cookit-be/features/steps"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		stepEntity steps.StepEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		result        *steps.StepEntity
		errResult     error
	}
}

func InsertStepTestTable() []TestTable {
	tname := "test insert step "
	return []TestTable{
		{
			Name: tname + "expect failed - empty step name",
			Input: struct {
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
				errResult: nil,
			},
		},
	}
}

func UpdateStepByIdTestTable() []TestTable {
	tname := "test update step by id "
	return []TestTable{
		{
			Name: tname + "expect failed - empty step name",
			Input: struct {
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
				errResult: nil,
			},
		},
	}
}

func DeleteStepByIdTestTable() []TestTable {
	tname := "test delete step by id"
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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

func DeleteStepByRecipeIdTestTable() []TestTable {
	tname := "test delete step by id"
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
				stepEntity steps.StepEntity
			}{
				stepEntity: steps.StepEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Name:     "Masukan air ke dalam wadah sebanyak 100mL",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *steps.StepEntity
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
