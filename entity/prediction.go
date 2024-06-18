package entity

import "time"

type Prediction struct {
	Pr_ID          string    `json:"id" gorm:"primaryKey"`
	Gambar         string    `json:"gambar" binding:"required"`
	Hasil_Prediksi string    `json:"hasil_prediksi" binding:"required"`
	Confidence	   float64    `json:"confidence" binding:"required"`
	Tgl            time.Time `json:"tgl" binding:"required"`
	UserID         string    `gorm:"foreignKey" json:"user_id"`
	DiseaseID 	   uint64    `gorm:"foreignKey" json:"disease_id"`

	User           *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Disease        *Disease  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"disease,omitempty"`
}
