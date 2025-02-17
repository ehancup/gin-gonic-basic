package guard

import (
	"fmt"
	"gin-gorm/config"
	"gin-gorm/src/utils/handler"
	"gin-gorm/src/utils/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthGuard() gin.HandlerFunc {

	return func(c *gin.Context) {
		logger.Info("im guarding auth")
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(handler.Throw401("Unauthorized"))
			return
		}
		splitAuth := strings.Split(auth, " ")
		if len(splitAuth) != 2 || splitAuth[0] != "Bearer" || splitAuth[1] == "" {
			c.AbortWithStatusJSON(handler.Throw401("Unauthorized"))
			return
		}
		tokenString := splitAuth[1]
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(config.GetConfig().App.JwtSecret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(handler.Throw401(err.Error()))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(handler.Throw401("Token Exp"))
				return
			}
			
			c.Set("user", claims)
			c.Next()
			
		} else {
			c.AbortWithStatusJSON(handler.Throw401("Invalid claims"))
		}
	}
}
