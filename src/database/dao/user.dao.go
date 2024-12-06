package dao

import (
	"time"

	"gorm.io/gorm"
)

// Entity

type UserEntity struct {
	gorm.Model
	Name     string    `gorm:"size:255;uniqueIndex;not null;"`
	Email    string    `gorm:"size:255;not null;"`
	Address  string    `gorm:"size:255;not null;"`
	BornDate time.Time `gorm:"column:born_date;not null;"`
}

func (UserEntity) TableName() string {
	return "users"
}