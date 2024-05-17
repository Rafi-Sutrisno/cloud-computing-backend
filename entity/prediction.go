package entity

type Prediction struct {
	ID             uint64 `json:"id" gorm:"primaryKey"`
	Gambar         string `json:"gambar" binding:"required"`
	Hasil_Prediksi string `json:"hasil_prediksi" binding:"required"`
	Tgl            string `json:"tgl" binding:"required"`
	UserID         uint64 `gorm:"foreignKey" json:"user_id"`
	User           *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}
