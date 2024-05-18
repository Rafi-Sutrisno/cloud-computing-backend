package repository

import (
	"context"
	"errors"
	"mods/dto"
	"mods/entity"

	"gorm.io/gorm"
)

type userConnection struct {
	connection *gorm.DB
}

type UserRepository interface {
	// functional
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUser(ctx context.Context, updateDTO dto.UpdateUserDTO, userID string) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := db.connection.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (db *userConnection) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := db.connection.Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (db *userConnection) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var listUser []entity.User

	tx := db.connection.Find(&listUser)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return listUser, nil
}

func (db *userConnection) DeleteUser(ctx context.Context, id string) error {
	var user entity.User
	tx := db.connection.Where("U_Id = ?", id).Delete(&user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (uc *userConnection) UpdateUser(ctx context.Context, updateDTO dto.UpdateUserDTO, userID string) (entity.User, error) {
	var user entity.User

	getUser := uc.connection.Where("id = ?", userID).Take(&user)
	if getUser.Error != nil {
		return entity.User{}, errors.New("invalid user id")
	}

	if tx := uc.connection.Model(&user).Updates(map[string]interface{}{"name": updateDTO.Name, "notelp": updateDTO.Notelp}); tx != nil {
		return entity.User{}, errors.New("failed to update")
	}

	updatedUser := uc.connection.Where("id = ?", userID).Take(&user)
	_ = updatedUser

	return user, nil
}
