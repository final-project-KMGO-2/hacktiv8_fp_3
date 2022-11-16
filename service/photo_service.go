package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, photoCreate entity.PhotoCreate) (entity.Photo, error)
	GetPhotos(ctx context.Context) ([]entity.Photo, error)
	GetPhotoByID(ctx context.Context, photoID uint64) (entity.Photo, error)
	UpdatePhoto(ctx context.Context, photoUpdate entity.PhotoUpdate) (entity.Photo, error)
	DeletePhoto(ctx context.Context, photoID uint64) error
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(pr repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: pr,
	}
}

func (s *photoService) CreatePhoto(ctx context.Context, photoCreate entity.PhotoCreate) (entity.Photo, error) {
	photo := entity.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoCreate))
	if err != nil {
		return photo, err
	}

	res, err := s.photoRepository.CreatePhoto(ctx, photo)
	if err != nil {
		return photo, err
	}

	return res, nil
}

func (s *photoService) GetPhotos(ctx context.Context) ([]entity.Photo, error) {
	res, err := s.photoRepository.GetPhotos(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *photoService) GetPhotoByID(ctx context.Context, photoID uint64) (entity.Photo, error) {
	res, err := s.photoRepository.GetPhotoByID(ctx, photoID)
	if err != nil {
		return entity.Photo{}, err
	}

	return res, nil
}

func (s *photoService) UpdatePhoto(ctx context.Context, photoUpdate entity.PhotoUpdate) (entity.Photo, error) {
	photo := entity.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoUpdate))
	if err != nil {
		return photo, err
	}

	res, err := s.photoRepository.UpdatePhoto(ctx, photo)
	if err != nil {
		return photo, err
	}

	return res, nil
}

func (s *photoService) DeletePhoto(ctx context.Context, photoID uint64) error {
	err := s.photoRepository.DeletePhoto(ctx, photoID)
	if err != nil {
		return err
	}

	return nil
}
