package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
)

type chatroomService struct {
	chatroomRepository repository.ChatroomRepository
}

type ChatRoomService interface {
	// functional
	CreateChatroom(ctx context.Context, chatroomDTO dto.CreateChatRoomDTO) (entity.ChatRoom, error)
	RemoveChatroom(ctx context.Context, id uint64) error
	GetChatroom(ctx context.Context, getchatroomDTO dto.GetChatRoomDTO) ([]entity.ChatRoom, error)
}

func NewChatRoomService(cr repository.ChatroomRepository) ChatRoomService {
	return &chatroomService{
		chatroomRepository: cr,
	}
}

func (cs *chatroomService) CreateChatroom(ctx context.Context, chatroomDTO dto.CreateChatRoomDTO) (entity.ChatRoom, error) {

	newChatroom := entity.ChatRoom{
		Uid: chatroomDTO.Uid,
		Uid_Doctor: chatroomDTO.Uid_Doctor,
	}

	return cs.chatroomRepository.AddChatroom(ctx, newChatroom)
}

func (cs *chatroomService) RemoveChatroom(ctx context.Context, id uint64) error {
	return cs.chatroomRepository.RemoveChatroom(ctx, id)
}

func (cs *chatroomService) GetChatroom(ctx context.Context, getchatroomDTO dto.GetChatRoomDTO) ([]entity.ChatRoom, error) {

	if(getchatroomDTO.Role == "User"){
		return cs.chatroomRepository.GetChatroomUser(ctx, getchatroomDTO.Test)
	}else {
		return cs.chatroomRepository.GetChatroomDoctor(ctx, getchatroomDTO.Test)
	}
	
}
