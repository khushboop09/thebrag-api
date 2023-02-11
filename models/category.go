package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

func (c *Category) TableName() string {
	return "categories"
}
