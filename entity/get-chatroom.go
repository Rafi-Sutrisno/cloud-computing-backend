package entity

type GetChatRoom struct {
	ID         uint64 `json:"id" gorm:"primaryKey"`
	Uid        string `json:"uid" binding:"required"`
	U_Name     string `json:"u_name" binding:"required"`
	Uid_Doctor string `json:"uid_doctor" binding:"required"`
	U_Doctor   string `json:"u_doctor" binding:"required"`
	Role       string `json:"role" binding:"required"`
}
