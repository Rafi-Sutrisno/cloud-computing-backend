package entity

type ChatRoom struct {
	ID         uint64 `json:"id" gorm:"primaryKey"`
	Uid        string `json:"uid" binding:"required"`
	Uid_Doctor string `json:"uid_doctor" binding:"required"`
}
