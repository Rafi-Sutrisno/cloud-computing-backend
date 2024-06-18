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
	GetAllDoctor(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUserName(ctx context.Context, updateDTO dto.UpdateNameUserDTO, userID string) (entity.User, error)
	UpdateUserNotelp(ctx context.Context, updateDTO dto.UpdateNotelpUserDTO, userID string) (entity.User, error)
	Me(ctx context.Context, id string) (entity.User, error)
	ProfilePicture(ctx context.Context, link string, uid string) (string, error)
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

func (db *userConnection) GetAllDoctor(ctx context.Context) ([]entity.User, error) {
	var listDoctor []entity.User
	var doctor = "Doctor"

	tx := db.connection.Where("role = ?", doctor).Find(&listDoctor)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return listDoctor, nil
}

func (db *userConnection) DeleteUser(ctx context.Context, id string) error {
	var user entity.User
	tx := db.connection.Where("U_Id = ?", id).Delete(&user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (uc *userConnection) UpdateUserName(ctx context.Context, updateDTO dto.UpdateNameUserDTO, userID string) (entity.User, error) {
	var user entity.User

	getUser := uc.connection.Where("u_id = ?", userID).Take(&user)
	if getUser.Error != nil {
		return entity.User{}, errors.New("invalid user id")
	}

	if tx := uc.connection.Model(&user).Update("name", updateDTO.Name).Error; tx != nil {
		return entity.User{}, tx
	}

	return user, nil
}

func (uc *userConnection) UpdateUserNotelp(ctx context.Context, updateDTO dto.UpdateNotelpUserDTO, userID string) (entity.User, error) {
	var user entity.User

	getUser := uc.connection.Where("u_id = ?", userID).Take(&user)
	if getUser.Error != nil {
		return entity.User{}, errors.New("invalid user id")
	}

	if tx := uc.connection.Model(&user).Update("notelp", updateDTO.Notelp).Error; tx != nil {
		return entity.User{}, tx
	}

	return user, nil
}

func (uc *userConnection) Me(ctx context.Context, id string) (entity.User, error){
	var user entity.User

	getUser := uc.connection.Where("u_id = ?", id).Preload("Prediction").Take(&user)
	if getUser.Error != nil {
		return entity.User{}, errors.New("invalid user id")
	}

	return user, nil
}

func(uc *userConnection) ProfilePicture(ctx context.Context, link string, uid string) (string, error) {
	var user entity.User

	getUser := uc.connection.Where("u_id = ?", uid).Take(&user)
	if getUser.Error != nil {
		return "failed to get user with the id", errors.New("invalid user id")
	}

	if tx := uc.connection.Model(&user).Updates(map[string]interface{}{"picture": link}).Error; tx != nil {
		return "failed to update user", tx
	}

	return "success to update profile picture.", nil
}