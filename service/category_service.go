package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, categoryCreate entity.CategoryCreate) (entity.Category, error)
	GetCategory(ctx context.Context) ([]entity.Category, error)
	PatchCategory(ctx context.Context, categoryID uint64, CategoryPatch entity.CategoryPatch) (entity.Category, error)
	DeleteCategory(ctx context.Context, categoryID uint64) error
}

type categoryService struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryRepository: cr,
	}
}

// CreateCategory implements CategoryService
func (s *categoryService) CreateCategory(ctx context.Context, categoryCreate entity.CategoryCreate) (entity.Category, error) {
	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&categoryCreate))
	if err != nil {
		return category, err
	}

	res, err := s.CategoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return category, err
	}
	return res, nil
}

// GetCategory implements CategoryService
func (s *categoryService) GetCategory(ctx context.Context) ([]entity.Category, error) {
	res, err := s.CategoryRepository.GetCategory(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PatchCategory implements CategoryService
func (s *categoryService) PatchCategory(ctx context.Context, categoryID uint64, CategoryPatch entity.CategoryPatch) (entity.Category, error) {
	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&CategoryPatch))
	if err != nil {
		return category, err
	}
	res, err := s.CategoryRepository.PatchCategory(ctx, category)
	if err != nil {
		return entity.Category{}, err
	}
	return res, nil
}

// DeleteCategory implements CategoryService
func (s *categoryService) DeleteCategory(ctx context.Context, categoryID uint64) error {
	err := s.CategoryRepository.DeleteCategory(ctx, categoryID)
	if err != nil {
		return err
	}
	return nil
}
