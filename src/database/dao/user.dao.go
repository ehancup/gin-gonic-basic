package dao

import (
	"time"

	"gorm.io/gorm"
)

// Entity

type UserEntity struct {
	ID       *uint     `gorm:"primarykey"`
	gorm.Model
	Name     string    `gorm:"size:255;not null;"`
	Email    string    `gorm:"size:255;not null;uniqueIndex"`
	Address  string    `gorm:"size:255;not null;"`
	BornDate time.Time `gorm:"column:born_date;not null;"`
}

func (UserEntity) TableName() string {
	return "users"
}

