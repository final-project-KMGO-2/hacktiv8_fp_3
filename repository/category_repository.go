package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	GetCategory(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	PatchCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	DeleteCategory(ctx context.Context, categoryID uint64) error
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

// CreateCategory implements CategoryRepository
func (db *categoryConnection) CreateCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	tx := db.connection.Create(&category)
	if tx.Error != nil {
		return entity.Category{}, tx.Error
	}
	return category, nil
}

// GetCategory implements CategoryRepository
func (db *categoryConnection) GetCategory(ctx context.Context) ([]entity.Category, error) {
	var category []entity.Category
	tx := db.connection.Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}

func (db *categoryConnection) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var category entity.Category
	tx := db.connection.Where(("id = ?"), id).Take(&category)
	if tx.Error != nil {
		return entity.Category{}, tx.Error
	}
	return category, nil
}

// PatchCategory implements CategoryRepository
func (db *categoryConnection) PatchCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	tx := db.connection.Save(&category)
	if tx.Error != nil {
		return entity.Category{}, tx.Error
	}
	return category, nil
}

// DeleteCategory implements CategoryRepository
func (db *categoryConnection) DeleteCategory(ctx context.Context, categoryID uint64) error {
	tx := db.connection.Delete(&entity.Category{}, categoryID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
