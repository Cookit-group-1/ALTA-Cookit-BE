package service

import (
	"alta-cookit-be/features/likes"
	"alta-cookit-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		likeEntity likes.LikeEntity
	}
	Output struct {
		IsError   bool
		errResult error
	}
}

func LikeRecipeTestTable() []TestTable {
	tname := "test like recipe"
	return []TestTable{
		{
			Name: tname + "expect failed - already liked",
			Input: struct {
				likeEntity likes.LikeEntity
			}{
				likeEntity: likes.LikeEntity{
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				IsError   bool
				errResult error
			}{
				IsError:   true,
				errResult: errors.New(consts.LIKE_AlreadyLiked),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				likeEntity likes.LikeEntity
			}{
				likeEntity: likes.LikeEntity{
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				IsError   bool
				errResult error
			}{
				IsError:   false,
				errResult: nil,
			},
		},
	}
}

func UnlikeRecipeTestTable() []TestTable {
	tname := "test like recipe"
	return []TestTable{
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				likeEntity likes.LikeEntity
			}{
				likeEntity: likes.LikeEntity{
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				IsError   bool
				errResult error
			}{
				IsError:   true,
				errResult: errors.New(consts.SERVER_ForbiddenRequest),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				likeEntity likes.LikeEntity
			}{
				likeEntity: likes.LikeEntity{
					UserID:   1,
					RecipeID: 1,
				},
			},
			Output: struct {
				IsError   bool
				errResult error
			}{
				IsError:   false,
				errResult: nil,
			},
		},
	}
}
