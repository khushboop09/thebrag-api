package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func (b *User) TableName() string {
	return "users"
}
