package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"column:user_name;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}
