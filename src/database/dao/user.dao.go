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
	Book []BookEntity `gorm:"foreignKey:user_id"` // nama "Book" yang akan dijadikan parameter untuk relasi
}

func (UserEntity) TableName() string {
	return "users"
}

