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
