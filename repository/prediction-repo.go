package repository

import (
	"context"
	"mods/entity"

	"gorm.io/gorm"
)

type predictionConnection struct {
	connection *gorm.DB
}

type PredictionRepository interface {
	AddPrediction(ctx context.Context, prediction entity.Prediction) (entity.Prediction, error)
	GetPredictionByUserID(ctx context.Context, UserID string) ([]entity.Prediction, error)
	GetPredictionByPredictionID(ctx context.Context, PredicitonID string) (entity.Prediction, error)
}

func NewPredictionRepository(db *gorm.DB) PredictionRepository {
	return &predictionConnection{
		connection: db,
	}
}

func (pc *predictionConnection) AddPrediction(ctx context.Context, prediction entity.Prediction) (entity.Prediction, error) {
	if err := pc.connection.Create(&prediction).Error; err != nil {
		return entity.Prediction{}, err
	}
	return prediction, nil
}

func (pc *predictionConnection) GetPredictionByUserID(ctx context.Context, UserID string) ([]entity.Prediction, error) {
	var prediciton []entity.Prediction

	if err := pc.connection.Where("user_id = ?", UserID).Find(&prediciton).Error; err != nil {
		return nil, err
	}

	return prediciton, nil
}

func (pc *predictionConnection) GetPredictionByPredictionID(ctx context.Context, PredicitonID string) (entity.Prediction, error) {
	var prediciton entity.Prediction

	if err := pc.connection.Where("Pr_ID = ?", PredicitonID).Take(&prediciton).Error; err != nil {
		return entity.Prediction{}, err
	}

	return prediciton, nil
}
