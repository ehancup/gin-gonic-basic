package users

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.RouterGroup, service *Service) {
	router := app.Group("/user")

	router.GET("/list", service.GetAllUser)
	router.GET("/detail/:id", service.GetById)
	router.POST("/create", service.CreateUser)
	router.PUT("/update/:id", service.UpdateById)
	router.DELETE("/delete/:id", service.DeleteById)
}