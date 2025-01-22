package dao

import (
	"gorm.io/gorm"
)

type BookEntity struct {
	ID   *uint  `gorm:"primarykey"`
	Name string `gorm:"size:255;not null;"`
	gorm.Model
}

func (BookEntity) TableName() string {
	return "book"
}
