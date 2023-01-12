package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Username     string
	TTRansaction []Transaction
	IsAdmin      bool
}

type UserResponse struct {
	gorm.Model
	Name        string
	Email       string
	Username    string
	IsAdmin     bool
	Transaction []Transaction
}

func (UserResponse) TableName() string { return "users" }
