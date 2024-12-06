package users

import (
	"errors"
	"gin-gorm/src/database/dao"
	"gin-gorm/src/utils/gldtr"
	"time"
)

// Response

type UserListResp struct {
	ID       *uint      `json:"id"`
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Address  *string    `json:"address"`
	BornDate *time.Time `json:"born_date"`
}

type UserDetailResp struct {
	ID       *uint      `json:"id"`
	Name     *string    `json:"name"`
	Email    *string    `json:"email"`
	Address  *string    `json:"address"`
	BornDate *time.Time `json:"born_date"`
}

// Request

type UserCreateReq struct {
	Name     string `json:"name" g:"required,min=3"`
	Email    string `json:"email" g:"required,email"`
	Address  string `json:"address" g:"required"`
	BornDate string `json:"born_date" g:"required"`
}

func (v UserCreateReq) ToEntity() (dao.UserEntity, error) {

	parsedTime, err := time.Parse("2006-01-02", v.BornDate)
	if err != nil {
		return dao.UserEntity{}, errors.New("time must be in format 'YYYY-MM-DD'")
	}
	return dao.UserEntity{
		Name:     v.Name,
		Address:  v.Address,
		Email:    v.Email,
		BornDate: parsedTime,
	}, nil
}

// Validation (DTO)

var UserCreateDto = gldtr.G.Validator(UserCreateReq{})
