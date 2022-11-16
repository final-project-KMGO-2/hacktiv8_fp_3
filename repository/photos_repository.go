package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(ctx context.Context, photo entity.Photo) (entity.Photo, error)
	GetPhotos(ctx context.Context) ([]entity.Photo, error)
	GetPhotoByID(ctx context.Context, photoID uint64) (entity.Photo, error)
	UpdatePhoto(ctx context.Context, photo entity.Photo) (entity.Photo, error)
	DeletePhoto(ctx context.Context, photoID uint64) error
}

type photoConnection struct {
	connection *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: db,
	}
}
func (db *photoConnection) CreatePhoto(ctx context.Context, photo entity.Photo) (entity.Photo, error) {
	tx := db.connection.Create(&photo)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	return photo, nil
}

func (db *photoConnection) GetPhotos(ctx context.Context) ([]entity.Photo, error) {
	var photos []entity.Photo
	tx := db.connection.Preload("User").Find(&photos)
	if tx.Error != nil {
		return []entity.Photo{}, tx.Error
	}

	return photos, nil
}

func (db *photoConnection) GetPhotoByID(ctx context.Context, photoID uint64) (entity.Photo, error) {
	var photo entity.Photo
	tx := db.connection.Where(("id = ?"), photoID).Take(&photo)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	return photo, nil
}

func (db *photoConnection) UpdatePhoto(ctx context.Context, photo entity.Photo) (entity.Photo, error) {
	tx := db.connection.Save(&photo)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	return photo, nil
}

func (db *photoConnection) DeletePhoto(ctx context.Context, photoID uint64) error {
	tx := db.connection.Delete(&entity.Photo{}, photoID)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
