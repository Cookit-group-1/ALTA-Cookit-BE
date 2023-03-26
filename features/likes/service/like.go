package service

import (
	"alta-cookit-be/features/likes"
)

type RecipeService struct {
	likeData likes.LikeData_
}

func New(likeData likes.LikeData_) likes.LikeService_ {
	return &RecipeService{
		likeData: likeData,
	}
}

func (s *RecipeService) LikeRecipe(entity *likes.LikeEntity) error {
	err := s.likeData.LikeRecipe(entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipeService) UnlikeRecipe(entity *likes.LikeEntity) error {
	err := s.likeData.UnlikeRecipe(entity)
	if err != nil {
		return err
	}
	return nil
}
