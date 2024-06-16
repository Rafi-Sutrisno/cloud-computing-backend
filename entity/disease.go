package entity

type Disease struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Headline    string `json:"headline" binding:"required"`
	Description string `json:"description" binding:"required"`
}
