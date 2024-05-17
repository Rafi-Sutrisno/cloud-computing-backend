package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"path/filepath"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	// functional
	CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (entity.User, error) {
	// filename := filepath.Base(userDTO.File.Filename)
	newUser := entity.User{
		Name:    userDTO.User.Name,
		Email:   userDTO.User.Email,
		Profile: filepath.Base(userDTO.Profile.Filename),
	}

	return us.userRepository.AddUser(ctx, newUser)
}

func (us *userService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	return us.userRepository.GetAllUser(ctx)
}

func (us *userService) DeleteUser(ctx context.Context, id uint64) error {
	return us.userRepository.DeleteUser(ctx, id)
}
