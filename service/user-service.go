package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	// functional
	CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (entity.User, error) {
	newUser := entity.User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
		// Profile: userDTO.Profile,
	}

	return us.userRepository.AddUser(ctx, newUser)
}

func (us *userService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	return us.userRepository.GetAllUser(ctx)
}
