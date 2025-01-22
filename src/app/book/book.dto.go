package book

import (
	"gin-gorm/src/database/dao"
	"gin-gorm/src/utils/gldtr"
)

// REQUEST

type BookReq struct {
	Name string `json:"name" validate:"required,min=3"`
}

type BookCreateReq struct {
	BookReq
}

func (v BookCreateReq) ToEntity() (dao.BookEntity, error) {
	return dao.BookEntity{
		Name: v.Name,
	},nil
}


// Validator (DTO)

var BookCreateDto = gldtr.G.Validator(BookCreateReq{})
