package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type SocmedService interface {
	GetSocmedInfo(ctx context.Context) ([]entity.SocialMedia, error)
	GetSocmedByID(ctx context.Context, socmedID uint64) (entity.SocialMedia, error)
	AddNewSocmed(ctx context.Context, socmedCreate entity.SocialMediaCreate) (entity.SocialMedia, error)
	DeleteSocmed(ctx context.Context, id uint64) error
	UpdateSocmed(ctx context.Context, socmedUpdate entity.SocialMediaUpdate) (entity.SocialMedia, error)
}

type socmedService struct {
	socmedRepository repository.SocmedRepository
	userRepository   repository.UserRepository
}

func NewSocmedService(sr repository.SocmedRepository, ur repository.UserRepository) SocmedService {
	return &socmedService{
		socmedRepository: sr,
		userRepository:   ur,
	}
}

func (sr *socmedService) GetSocmedInfo(ctx context.Context) ([]entity.SocialMedia, error) {
	res, err := sr.socmedRepository.GetSocmeds(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sr *socmedService) GetSocmedByID(ctx context.Context, socmedID uint64) (entity.SocialMedia, error) {
	res, err := sr.socmedRepository.GetSocmedByID(ctx, socmedID)
	if err != nil {
		return entity.SocialMedia{}, err
	}
	return res, nil
}

func (sr *socmedService) AddNewSocmed(ctx context.Context, socmedCreate entity.SocialMediaCreate) (entity.SocialMedia, error) {
	socmed := entity.SocialMedia{}
	err := smapping.FillStruct(&socmed, smapping.MapFields(&socmedCreate))
	if err != nil {
		return entity.SocialMedia{}, err
	}

	res, err := sr.socmedRepository.CreateSocmed(ctx, socmed)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return res, nil
}
func (sr *socmedService) DeleteSocmed(ctx context.Context, id uint64) error {
	err := sr.socmedRepository.DeleteSocmed(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (sr *socmedService) UpdateSocmed(ctx context.Context, socmedUpdate entity.SocialMediaUpdate) (entity.SocialMedia, error) {
	socmed := entity.SocialMedia{}
	err := smapping.FillStruct(&socmed, smapping.MapFields(&socmedUpdate))
	if err != nil {
		return entity.SocialMedia{}, err
	}

	resp, err := sr.socmedRepository.UpdateSocmed(ctx, socmed)
	if err != nil {
		return entity.SocialMedia{}, err
	}
	return resp, nil
}
