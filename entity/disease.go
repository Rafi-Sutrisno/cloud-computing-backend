package entity

import "github.com/google/uuid"

type Disease struct {
	D_Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
}
