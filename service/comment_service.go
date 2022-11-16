package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type CommentService interface {
	CreateComment(ctx context.Context, commentCreate entity.CommentCreate) (entity.Comment, error)
	GetComment(ctx context.Context, userID uint64) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error)
	UpdateCommentByID(ctx context.Context, commentID uint64, commentUpdate entity.CommentUpdate) (entity.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID uint64) error
}

type commentService struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(cr repository.CommentRepository) CommentService {
	return &commentService{
		CommentRepository: cr,
	}
}

func (s *commentService) CreateComment(ctx context.Context, commentCreate entity.CommentCreate) (entity.Comment, error) {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&commentCreate))
	if err != nil {
		return comment, err
	}

	res, err := s.CommentRepository.CreateComment(ctx, comment)
	if err != nil {
		return comment, err
	}

	return res, nil
}

func (s *commentService) GetComment(ctx context.Context, userID uint64) ([]entity.Comment, error) {
	res, err := s.CommentRepository.GetComment(ctx, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *commentService) GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error) {
	res, err := s.CommentRepository.GetCommentByID(ctx, commentID)
	if err != nil {
		return entity.Comment{}, err
	}

	return res, nil
}

func (s *commentService) UpdateCommentByID(ctx context.Context, commentID uint64, commentUpdate entity.CommentUpdate) (entity.Comment, error) {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&commentUpdate))
	if err != nil {
		return comment, err
	}
	res, err := s.CommentRepository.UpdateCommentByID(ctx, comment)
	if err != nil {
		return entity.Comment{}, err
	}

	return res, nil
}

func (s *commentService) DeleteCommentByID(ctx context.Context, commentID uint64) error {
	err := s.CommentRepository.DeleteCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	return nil
}
