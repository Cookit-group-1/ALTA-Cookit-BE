package service

import (
	"alta-cookit-be/features/comments"
	"alta-cookit-be/utils/consts"

	"errors"

	"github.com/go-playground/validator"
)

type CommentService struct {
	commentData comments.CommentData_
	validate       *validator.Validate
}

func New(commentData comments.CommentData_) comments.CommentService_ {
	return &CommentService{
		commentData: commentData,
		validate:       validator.New(),
	}
}

func (s *CommentService) InsertComment(entity *comments.CommentEntity) (*comments.CommentEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	output, err := s.commentData.InsertComment(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *CommentService) UpdateCommentById(entity *comments.CommentEntity) (*comments.CommentEntity, error) {
	err := s.validate.Struct(entity)
	if err != nil {
		return nil, errors.New(consts.VALIDATION_InvalidInput)
	}

	output, err := s.commentData.UpdateCommentById(entity)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *CommentService) DeleteCommentById(entity *comments.CommentEntity) error {
	err := s.validate.Struct(entity)
	if err != nil {
		return errors.New(consts.VALIDATION_InvalidInput)
	}

	err = s.commentData.DeleteCommentById(entity)
	if err != nil {
		return err
	}
	return nil
}
