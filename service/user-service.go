package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"

	"github.com/google/uuid"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	// functional
	CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, userDTO dto.CreateUserDTO) (entity.User, error) {

	id := uuid.NewString()

	newUser := entity.User{
		U_Id:   id,
		Name:   userDTO.Name,
		Email:  userDTO.Email,
		Notelp: userDTO.Notelp,
		Pass:   userDTO.Pass,
		Role:   "User",
	}

	return us.userRepository.AddUser(ctx, newUser)
}

func (us *userService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	return us.userRepository.GetAllUser(ctx)
}

func (us *userService) DeleteUser(ctx context.Context, id string) error {
	return us.userRepository.DeleteUser(ctx, id)
}
