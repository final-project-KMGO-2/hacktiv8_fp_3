package entity

import (
	"hacktiv8_fp_2/helpers"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	BaseModel
}

type UserRegister struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type UserUpdate struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashAndSalt(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashAndSalt(u.Password)
	if err != nil {
		return err
	}
	return nil
}
