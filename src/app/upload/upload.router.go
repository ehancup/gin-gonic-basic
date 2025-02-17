package upload

import (
	"fmt"
	"gin-gorm/src/utils/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 	UploadFile godoc
//
//	@Summary		Upload Single File
//	@Description	Upload Single File with formData
//	@Tags			upload
//	@Accept			mpfd
//	@Produce		json
//	@Param			file			formData	file	true	"file will be uploaded"
//	@Success		200				{object}	map[string]string
//	@Failure		400				{object}	map[string]string
//	@Router			/upload/single	[post]
func singleUpload(ctx *gin.Context) {
	file, fileErr := ctx.FormFile("file")
	
	if fileErr != nil {
		ctx.JSON(handler.Throw500(fileErr.Error()))
		return 
	}

	if _, err := os.Stat("public"); os.IsNotExist(err) {
		os.Mkdir("public", os.ModePerm)
	}

	filePath := fmt.Sprintf("public/%s", file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(handler.Throw500("Failed to save file"))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully", "filename": fmt.Sprintf("http://localhost:3010/public/%s", file.Filename),})
}

func InitRoutes(app *gin.RouterGroup) {
	router := app.Group("/upload")

	router.POST("/single", singleUpload)
}
