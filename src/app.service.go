package src

import (
	"gin-gorm/src/app/auth"
	"gin-gorm/src/app/book"
	"gin-gorm/src/app/users"
)

func GetUserService() *users.Service {
	return &users.Service{}
}

func GetBookService() *book.Service {
	return &book.Service{}
}

func GetAuthService() *auth.Service {
	return &auth.Service{}
}