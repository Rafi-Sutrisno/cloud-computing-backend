package repository

import (
	"context"
	"fmt"
	"mods/entity"

	"gorm.io/gorm"
)

type diseaseConnection struct {
	connection *gorm.DB
}

type DiseaseRepository interface {
	// functional
	AddDisease(ctx context.Context, disease entity.Disease) (entity.Disease, error)
	GetAllDisease(ctx context.Context) ([]entity.Disease, error)
	DeleteDisease(ctx context.Context, id uint64) error
	GetDiseaseByID(ctx context.Context, id uint64) (entity.Disease, error)
}

func NewDiseaseRepository(db *gorm.DB) DiseaseRepository {
	return &diseaseConnection{
		connection: db,
	}
}

func (db *diseaseConnection) AddDisease(ctx context.Context, disease entity.Disease) (entity.Disease, error) {
	if err := db.connection.Create(&disease).Error; err != nil {
		return entity.Disease{}, err
	}

	return disease, nil
}

func (db *diseaseConnection) GetAllDisease(ctx context.Context) ([]entity.Disease, error) {
	var listDisease []entity.Disease

	tx := db.connection.Order("id ASC").Find(&listDisease)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return listDisease, nil
}

func (db *diseaseConnection) DeleteDisease(ctx context.Context, id uint64) error {
	var disease entity.Disease
	tx := db.connection.Where("id = ?", id).Delete(&disease)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		err := fmt.Errorf("no record found with id: %v", id)
		return err
	}

	return nil
}

func (dc *diseaseConnection) GetDiseaseByID(ctx context.Context, id uint64) (entity.Disease, error) {
	var disease entity.Disease
	if err := dc.connection.Where("id = ?", id).Take(&disease).Error; err != nil {
		return entity.Disease{}, err
	}
	return disease, nil
}
