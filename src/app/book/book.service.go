package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {}

func (Service) SayHello(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message" : "hello world!",
	})
}