package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func (c *User) TableName() string {
	return "users"
}
