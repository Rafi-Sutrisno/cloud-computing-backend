package entity

import (
	"mods/utils"

	"gorm.io/gorm"
)

type User struct {
	U_Id   string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Notelp string `json:"notelp" binding:"required"`
	Pass   string `json:"pass" binding:"required"`
	Role   string `json:"role" binding:"required"`
	Picture string `json:"picture"`

	Prediction []Prediction `json:"Prediciton,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Pass, err = utils.PasswordHash(u.Pass)
	if err != nil {
		return err
	}
	return nil
}
