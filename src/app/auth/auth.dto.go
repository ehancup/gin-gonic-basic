package auth

import "gin-gorm/src/database/dao"

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (v RegisterReq) ToEntity(hash []byte) dao.AuthEntity {
	return dao.AuthEntity{
		Email: v.Email,
		Password: string(hash),
	}
}