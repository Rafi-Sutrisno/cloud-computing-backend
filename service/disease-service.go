package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
)

type diseaseService struct {
	diseaseRepository repository.DiseaseRepository
}

type DiseaseService interface {
	// functional
	CreateDisease(ctx context.Context, diseaseDTO dto.CreateDiseaseDTO) (entity.Disease, error)
	GetAllDisease(ctx context.Context) ([]entity.Disease, error)
	DeleteDisease(ctx context.Context, id uint64) error
}

func NewDiseaseService(dr repository.DiseaseRepository) DiseaseService {
	return &diseaseService{
		diseaseRepository: dr,
	}
}

func (ds *diseaseService) CreateDisease(ctx context.Context, diseaseDTO dto.CreateDiseaseDTO) (entity.Disease, error) {

	newDesease := entity.Disease{
		Name:        diseaseDTO.Name,
		Description: diseaseDTO.Description,
	}

	return ds.diseaseRepository.AddDisease(ctx, newDesease)
}

func (ds *diseaseService) GetAllDisease(ctx context.Context) ([]entity.Disease, error) {
	return ds.diseaseRepository.GetAllDisease(ctx)
}

func (ds *diseaseService) DeleteDisease(ctx context.Context, id uint64) error {
	return ds.diseaseRepository.DeleteDisease(ctx, id)
}
