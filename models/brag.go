package models

import (
	"gorm.io/gorm"
)

type Brag struct {
	gorm.Model
	Title      string
	Details    string
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
	UserId     int
	User       User `gorm:"foreignKey:UserId"`
}

func (b *Brag) TableName() string {
	return "brags"
}
