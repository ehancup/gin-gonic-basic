package src

import (
	// "gin-gorm/app/service"

	"gin-gorm/src/app/auth"
	"gin-gorm/src/app/book"
	"gin-gorm/src/app/upload"
	"gin-gorm/src/app/users"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app.Group("/v1")
	

	users.InitRoutes(route, GetUserService())
	book.InitRoutes(route, GetBookService())
	auth.InitRoutes(route, GetAuthService())
	upload.InitRoutes(route)
}