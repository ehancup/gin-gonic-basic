package upload

import "github.com/gin-gonic/gin"

func InitRoutes(app *gin.RouterGroup) {
	router := app.Group("/upload")

	router.POST("/", func (ctx *gin.Context) {
		
	})
}