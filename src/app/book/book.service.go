package book

import (
	// "net/http"

	"gin-gorm/src/database"
	"gin-gorm/src/utils/handler"

	"github.com/gin-gonic/gin"
)

type Service struct {}

func (Service) CreateBook (ctx *gin.Context) {
	payload := handler.GetBody[BookCreateReq](ctx)
	if ctx.IsAborted() {
		return
	}
	
	bookEn, errEn := payload.ToEntity()
	if errEn != nil {
		ctx.JSON(handler.Throw500(errEn.Error()))
	}

	if err := database.DB.Table("book").Create(&bookEn).Error; err != nil {
		ctx.JSON(handler.Throw500(errEn.Error()))
	}

	ctx.JSON(200, gin.H{
		"message" : "success",
	})
}