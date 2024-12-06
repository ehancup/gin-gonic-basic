package src

import (
	// "gin-gorm/app/service"

	"gin-gorm/src/app/users"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app.Group("/v1")

	users.InitRoutes(route, GetUserService())
}