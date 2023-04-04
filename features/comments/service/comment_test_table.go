package service

import (
	"alta-cookit-be/features/comments"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		commentEntity comments.CommentEntity
	}
	Output struct {
		isErrValidate bool
		isEntitled    bool
		IsError       bool
		result        *comments.CommentEntity
		results       *[]comments.CommentEntity
		errResult     error
	}
}

func SelectCommentsByRecipeIdTestTable() []TestTable {
	tname := "test select comments by recipe id "
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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

func InsertCommentTestTable() []TestTable {
	tname := "test insert comment "
	return []TestTable{
		{
			Name: tname + "expect failed - empty comment",
			Input: struct {
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
			Name: tname + "expect success - filled comment",
			Input: struct {
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
				errResult: nil,
			},
		},
	}
}

func UpdateCommentByIdTestTable() []TestTable {
	tname := "test update comment by id "
	return []TestTable{
		{
			Name: tname + "expect failed - empty comment",
			Input: struct {
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
			Name: tname + "expect success - filled comment",
			Input: struct {
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
				errResult     error
			}{
				isErrValidate: false,
				isEntitled:    true,
				IsError:       false,
				result: &comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
				errResult: nil,
			},
		},
	}
}

func DeleteCommentByIdTestTable() []TestTable {
	tname := "test delete comment by id "
	return []TestTable{
		{
			Name: tname + "expect failed - is not entitled",
			Input: struct {
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
				commentEntity comments.CommentEntity
			}{
				commentEntity: comments.CommentEntity{
					ID:       1,
					UserID:   1,
					RecipeID: 1,
					Comment:  "Enak banget ini",
				},
			},
			Output: struct {
				isErrValidate bool
				isEntitled    bool
				IsError       bool
				result        *comments.CommentEntity
				results       *[]comments.CommentEntity
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
