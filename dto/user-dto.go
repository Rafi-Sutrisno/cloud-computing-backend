package dto

import "mime/multipart"

type CreateUserDTO struct {
	User struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	} `form:"data"`
	Profile *multipart.FileHeader `form:"file"`
}

type LoginDTO struct {
	Name string `json:"name" binding:"required"`
}

type PredictionDTO struct {
	Prediction int `json:"prediction"`
}
