package models

import (
	"gorm.io/gorm"
)

type Brag struct {
	gorm.Model
	Title   string
	Details string
}

func (b *Brag) TableName() string {
	return "brags"
}
