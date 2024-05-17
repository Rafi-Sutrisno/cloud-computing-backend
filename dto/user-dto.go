package dto

type CreateUserDTO struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Notelp string `json:"notelp" binding:"required"`
	Pass   string `json:"pass" binding:"required"`
}

type LoginDTO struct {
	Name string `json:"name" binding:"required"`
}

type PredictionDTO struct {
	Prediction int `json:"prediction"`
}
