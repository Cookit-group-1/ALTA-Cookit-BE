package service

import (
	"alta-cookit-be/features/images"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		imageEntities *[]images.ImageEntity
		imageEntity   *images.ImageEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		results       *[]images.ImageEntity
		result        *images.ImageEntity
		errResult     error
	}
}

func InsertImageTestTable() []TestTable {
	tname := "test insert image "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntities: &[]images.ImageEntity{
					images.ImageEntity{
						UserID:   1,
						RecipeID: 1,
					},
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntities: &[]images.ImageEntity{
					images.ImageEntity{
						UserID:   1,
						RecipeID: 1,
					},
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntities: &[]images.ImageEntity{
					images.ImageEntity{
						UserID:   1,
						RecipeID: 1,
					},
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntities: &[]images.ImageEntity{
					images.ImageEntity{
						UserID:   1,
						RecipeID: 1,
					},
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				results: &[]images.ImageEntity{
					images.ImageEntity{
						ID:       1,
						UserID:   1,
						RecipeID: 1,
						UrlImage: "https://test.com/",
					},
				},
				errResult: nil,
			},
		},
	}
}

func UpdateImageByIdTestTable() []TestTable {
	tname := "test update image by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					UrlImage: "https://test.com/",
				},
				errResult: nil,
			},
		},
	}
}

func DeleteImageByIdTestTable() []TestTable {
	tname := "test delete image by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					UrlImage: "https://test.com/",
				},
				errResult: nil,
			},
		},
	}
}

func DeleteImageByRecipeIdTestTable() []TestTable {
	tname := "test delete image by recipe id"
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    false,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success - is entitled",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     nil,
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       true,
				results:       &[]images.ImageEntity{},
				result:        &images.ImageEntity{},
				errResult:     errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				imageEntities *[]images.ImageEntity
				imageEntity   *images.ImageEntity
			}{
				imageEntity: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				results       *[]images.ImageEntity
				result        *images.ImageEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &images.ImageEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					UrlImage: "https://test.com/",
				},
				errResult: nil,
			},
		},
	}
}
