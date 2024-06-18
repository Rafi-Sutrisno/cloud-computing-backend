package service

import (
	"context"
	"fmt"
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"mods/utils"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type predictionService struct {
	predictionRepository repository.PredictionRepository
	diseaseRepository repository.DiseaseRepository
}

type PredictionService interface {
	CreatePrediction(ctx context.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error)
	GetPredictionByUserID(ctx context.Context, UserID string) ([]entity.Prediction, error)
	GetPredictionByPredictionID(ctx context.Context, PredictionID string) (entity.Prediction, error)
	DeletePredictionbyId(ctx context.Context, PredictionID string, PredictLink string) ( error)
}

func NewPredictionService(pr repository.PredictionRepository,  dr repository.DiseaseRepository) PredictionService {
	return &predictionService{
		predictionRepository: pr,
		diseaseRepository: dr,
	}
}

func (ps *predictionService) CreatePrediction(ctx context.Context, predictionDTO dto.PredictImageDTO, userID string) (entity.Prediction, error) {
	id := uuid.NewString()
	imageFile := predictionDTO.File

	img_uuid, err := utils.UploadToBucket(imageFile, "prediction")
	if err != nil {
		return entity.Prediction{}, err
	}

	temp := strings.Split(img_uuid, "/")
	img_name := temp[len(temp)-1]

	result, confidence, err := utils.PredictionAPI(img_name)
	if err != nil {
		return entity.Prediction{}, err
	}

	result_int, _ := strconv.ParseUint(result, 10, 64)

	confidence_float, _ := strconv.ParseFloat(confidence, 64)

	var prediksi entity.Disease 

	prediksi, _ = ps.diseaseRepository.GetDiseaseByID(ctx, result_int+1)

	link := "https://storage.googleapis.com/example-bucket-test-cc-trw/" + img_uuid

	newPrediction := entity.Prediction{
		Pr_ID:          id,
		Gambar:         link,
		Hasil_Prediksi: prediksi.Name,
		Confidence:     confidence_float,
		Tgl:            time.Now(),
		UserID:         userID,
		DiseaseID: 		result_int+1,
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