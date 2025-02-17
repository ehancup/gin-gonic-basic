package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Throw422(msg string) (int, gin.H) {
	return http.StatusUnprocessableEntity, gin.H{
		"message" : msg,
	}
}
func Throw404(msg string) (int, gin.H) {
	return http.StatusNotFound, gin.H{
		"message" : msg,
	}
}
func Throw500(msg string) (int, gin.H) {
	return http.StatusInternalServerError, gin.H{
		"message" : msg,
	}
}
func Throw401(msg string) (int, gin.H) {
	return http.StatusUnauthorized, gin.H{
		"message" : msg,
	}
}