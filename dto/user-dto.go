package dto

import "mime/multipart"

type CreateUserDTO struct {
	Name    string                `json:"name" binding:"required"`
	Email   string                `json:"email" binding:"required"`
	Profile *multipart.FileHeader `form:"file"`
}

type LoginDTO struct {
	Name string `json:"name" binding:"required"`
}

type PredictionDTO struct {
	Prediction int `json:"prediction"`
}
