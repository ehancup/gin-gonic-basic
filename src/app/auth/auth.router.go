package auth

import "github.com/gin-gonic/gin"

func InitRoutes(app *gin.RouterGroup, service *Service) {
	
	app.POST("/login", service.Login)
	app.POST("/register", service.Register)
}