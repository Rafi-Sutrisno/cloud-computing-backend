package entity

type User struct {
	U_Id   string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Notelp string `json:"notelp" binding:"required"`
	Pass   string `json:"pass" binding:"required"`
	Role   string `json:"role" binding:"required"`
}
