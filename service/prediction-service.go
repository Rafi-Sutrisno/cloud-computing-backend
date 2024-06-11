package service

import (
	"context"
	"fmt"
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"mods/utils"
	"time"

	"github.com/google/uuid"
)

type predictionService struct {
	predictionRepository repository.PredictionRepository
}

type PredictionService interface {
	CreatePrediction(ctx context.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error)
	GetPredictionByUserID(ctx context.Context, UserID string) ([]entity.Prediction, error)
	GetPredictionByPredictionID(ctx context.Context, PredictionID string) (entity.Prediction, error)
	DeletePredictionbyId(ctx context.Context, PredictionID string, PredictLink string) ( error)
}

func NewPredictionService(pr repository.PredictionRepository) PredictionService {
	return &predictionService{
		predictionRepository: pr,
	}
}

func (ps *predictionService) CreatePrediction(ctx context.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error) {
	id := uuid.NewString()
	imageFile := predictionDTO.File

	img_uuid, err := utils.UploadToBucket(imageFile, "prediction")
	if err != nil {
		return entity.Prediction{}, err
	}

	result, err := utils.PredictionAPI(img_uuid)
	if err != nil {
		return entity.Prediction{}, err
	}

	link := "https://storage.googleapis.com/example-bucket-test-cc-trw/" + img_uuid

	newPrediction := entity.Prediction{
		Pr_ID:          id,
		Gambar:         link,
		Hasil_Prediksi: result,
		Tgl:            time.Now(),
		UserID:         userID,
	}

	return ps.predictionRepository.AddPrediction(ctx, newPrediction)
}

func (ps *predictionService) GetPredictionByUserID(ctx context.Context, UserID string) ([]entity.Prediction, error) {
	return ps.predictionRepository.GetPredictionByUserID(ctx, UserID)
}

func (ps *predictionService) GetPredictionByPredictionID(ctx context.Context, PredictionID string) (entity.Prediction, error) {
	return ps.predictionRepository.GetPredictionByPredictionID(ctx, PredictionID)
}

func (ps *predictionService) DeletePredictionbyId(ctx context.Context, PredictionID string, PredictLink string) ( error) {
	fmt.Print("in service")
	err := utils.DeleteFromBucket("prediction", PredictLink)
	if err != nil {
		return err
	}

	return ps.predictionRepository.DeletePredictionbyId(ctx, PredictionID)
}