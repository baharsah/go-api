package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	Transaction []Transaction `gorm:"foreignKey:BillerID"`
	IsAdmin     bool
}

type UserResponse struct {
	gorm.Model
	Name        string
	Email       string
	Username    string
	IsAdmin     bool
	Transaction []Transaction `gorm:"foreignKey:BillerID"`
}

func (UserResponse) TableName() string { return "users" }
