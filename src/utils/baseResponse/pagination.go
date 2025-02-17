package br

import (
	// "gin-gorm/src/utils/logger"
	"math"
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginationResponse struct {
	Pagination
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

func (v Pagination) GetResponse(total int) PaginationResponse {
	totalPage := math.Ceil(float64(total) / float64(v.PageSize))	
	return PaginationResponse{
		Pagination: v,
		Total: total,
		TotalPage: int(totalPage),
	}
}