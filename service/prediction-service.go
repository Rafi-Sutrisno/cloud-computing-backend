package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"mods/utils"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type predictionService struct {
	predictionRepository repository.PredictionRepository
}

type PredictionService interface {
	CreatePrediction(ctx context.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error)
}

func NewPredictionService(pr repository.PredictionRepository) PredictionService {
	return &predictionService{
		predictionRepository: pr,
	}
}

func (ps *predictionService) CreatePrediction(ctx context.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error) {
	id := uuid.NewString()
	imageFile := predictionDTO.File

	err := utils.UploadToBucket(imageFile)
	if err != nil {
		return entity.Prediction{}, err
	}

	imageName := filepath.Base(imageFile.Filename)
	result, err := utils.PredictionAPI(imageName)
	if err != nil {
		return entity.Prediction{}, err
	}

	newPrediction := entity.Prediction{
		Pr_ID:          id,
		Gambar:         imageName,
		Hasil_Prediksi: result,
		Tgl:            time.Now(),
		UserID:         userID,
	}

	return ps.predictionRepository.AddPrediction(ctx, newPrediction)
}
