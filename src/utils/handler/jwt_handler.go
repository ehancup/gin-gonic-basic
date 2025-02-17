package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ReqAuth struct {
	Exp   int64  `json:"exp"`
	Iat   int64  `json:"iat"`
	Email string `json:"email"`
	Sub   string `json:"sub"`
}

func GetAuthFromToken(c *gin.Context) (ReqAuth, error) {
	raw, exist := c.Get("user")

	if !exist {
		return ReqAuth{}, errors.New("user claims not found")
	}

	claims, ok := raw.(jwt.MapClaims)
	if !ok {
		return ReqAuth{}, errors.New("invalid token")
	}

	return ReqAuth{
		Exp: int64(claims["exp"].(float64)),
		Iat: int64(claims["iat"].(float64)),
		Email: claims["email"].(string),
		Sub: claims["sub"].(string),
	}, nil
}