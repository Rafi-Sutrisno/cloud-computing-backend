package dto

type CreateChatRoomDTO struct {
	Uid        string `json:"uid" binding:"required"`
	Uid_Doctor string `json:"uid_doctor" binding:"required"`
}

type GetChatRoomDTO struct {
	Test string `json:"test" binding:"required"`
	Role string `json:"role" binding:"required"`
}