package dto

type CreateDiseaseDTO struct {
	Name        string `json:"name" binding:"required"`
	Headline    string `json:"headline" binding:"required"`
	Description string `json:"description" binding:"required"`
}
