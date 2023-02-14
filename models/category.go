package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string
	UserId int
	User   User `gorm:"foreignKey:UserId"`
}

func (c *Category) TableName() string {
	return "categories"
}
