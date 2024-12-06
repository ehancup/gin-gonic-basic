package src

import "gin-gorm/src/app/users"

func GetUserService() *users.Service {
	return &users.Service{}
}