package service

import (
	"context"
	"mods/dto"
	"mods/entity"
	"mods/repository"
	"mods/utils"

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
	IsDuplicateEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	VerifyCredential(ctx context.Context, email string, pass string) (bool, error)
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

func (us *userService) IsDuplicateEmail(ctx context.Context, email string) (bool, error) {
	checkUser, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if checkUser.Email == "" {
		return false, nil
	}

	return true, nil
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return us.userRepository.GetUserByEmail(ctx, email)
}

func (us *userService) GetAllUser(ctx context.Context) ([]entity.User, error) {
	return us.userRepository.GetAllUser(ctx)
}

func (us *userService) DeleteUser(ctx context.Context, id string) error {
	return us.userRepository.DeleteUser(ctx, id)
}

func (us *userService) VerifyCredential(ctx context.Context, email string, pass string) (bool, error) {
	checkUser, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	checkPassword, err := utils.PasswordCompare(checkUser.Pass, []byte(pass))
	if err != nil {
		return false, err
	}

	if checkUser.Email == email && checkPassword {
		return true, nil
	}
	return false, nil
}
