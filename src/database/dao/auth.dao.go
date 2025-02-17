package dao

type AuthEntity struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
}

func (AuthEntity) TableName() string {
	return "auth"
}