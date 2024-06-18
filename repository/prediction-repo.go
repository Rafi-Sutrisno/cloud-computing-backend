package repository

import (
	"context"
	"fmt"
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
	DeletePredictionbyId(ctx context.Context, PredicitonID string) (error)
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

func (pc *predictionConnection) DeletePredictionbyId(ctx context.Context, PredicitonID string) ( error) {
	var prediction entity.Prediction

	tx := pc.connection.Where("pr_id = ?", PredicitonID).Delete(&prediction)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		err := fmt.Errorf("no record found with id: %v", PredicitonID)
		return err
	}
	
	return nil
}