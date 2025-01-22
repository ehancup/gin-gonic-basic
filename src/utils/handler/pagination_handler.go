package handler

import (
	br "gin-gorm/src/utils/baseResponse"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPagination(ctx *gin.Context) (br.Pagination, error) {
	page := ctx.Query("page")
	pageSize := ctx.Query("page_size")

	if page == "" {
		page = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil {
		return br.Pagination{}, errPage
	}

	pageSizeInt, errPageSize := strconv.Atoi(pageSize)
	if errPageSize != nil  {
		return br.Pagination{}, errPageSize
	}

	if pageInt < 1 {
		pageInt = 1
	}

	pagination := br.Pagination{
		Page: pageInt,
		PageSize: pageSizeInt,
	}

	return pagination, nil
}