package book

import "github.com/gin-gonic/gin"

func InitRoutes(app *gin.RouterGroup, service *Service) {
	router := app.Group("/book")

	router.POST("/create", service.CreateBook)
}