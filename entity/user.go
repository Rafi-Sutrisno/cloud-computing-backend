package entity

type User struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	// Profile string `json:"profile" binding:"required"`
}
