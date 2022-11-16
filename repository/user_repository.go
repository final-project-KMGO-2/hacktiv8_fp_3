package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserById(ctx context.Context, id uint64) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, userID uint64) error
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	tx := db.connection.Create(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

func (db *userConnection) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	tx := db.connection.Where(("email = ?"), email).Take(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (db *userConnection) GetUserById(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User
	tx := db.connection.Where(("id = ?"), id).Take(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (db *userConnection) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	tx := db.connection.Where(("username = ?"), username).Take(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (db *userConnection) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	tx := db.connection.Save(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

func (db *userConnection) DeleteUser(ctx context.Context, userID uint64) error {
	tx := db.connection.Delete(&entity.User{}, userID)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
