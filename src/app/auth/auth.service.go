package auth

import (
	// "gin-gorm/config"
	"gin-gorm/config"
	"gin-gorm/src/database"
	"gin-gorm/src/database/dao"
	br "gin-gorm/src/utils/baseResponse"
	"gin-gorm/src/utils/handler"
	"gin-gorm/src/utils/logger"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

// Login godoc
//	@Summary		Login
//	@Description	Login with credential
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		LoginReq	true	"Login Payload"
//	@Success		200		{object}	br.BaseSuccessResponse
//	@Router			/login	[post]
func (Service) Login(ctx *gin.Context) {
	payload := handler.GetBody[LoginReq](ctx)
	if ctx.IsAborted() {
		return
	}

	var user dao.AuthEntity

	if err := database.DB.Where("email = ?", payload.Email).Find(&user).Error; err != nil {
		ctx.JSON(handler.Throw500(err.Error()))
		return
	}

	if user.ID == 0 {
		ctx.JSON(handler.Throw404("no auth found!"))

		return
	}
	logger.Debug("user", "data", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		ctx.JSON(handler.Throw422("wrong password!"))
		return
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, // pakai HS, jangan ES

		jwt.MapClaims{
			"sub":   strconv.Itoa(int(user.ID)),
			"email": user.Email,
			"iat":   time.Now().Unix(),
			"exp":   time.Now().Add(time.Minute * 2).Unix(),
		})
	token, errToken := t.SignedString([]byte(config.GetConfig().App.JwtSecret))
	if errToken != nil {

		ctx.JSON(handler.Throw500(errToken.Error()))
		return
	}

	ctx.JSON(200, br.BaseSuccessResponse{
		Message: "success login",
		Success: true,
		Data: gin.H{
			"email":        user.Email,
			"access_token": token,
			"user":         user,
		},
	})

}

// Register godoc
//	@Summary		Register
//	@Description	Register with credential
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload		body		RegisterReq	true	"Register Payload"
//	@Success		200			{object}	br.BaseSuccessResponse
//	@Router			/register	[post]
func (Service) Register(ctx *gin.Context) {
	payload := handler.GetBody[RegisterReq](ctx)
	if ctx.IsAborted() {
		return
	}

	hash, errHash := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)

	if errHash != nil {
		ctx.JSON(handler.Throw500(errHash.Error()))
		return
	}

	regEntity := payload.ToEntity(hash)

	if err := database.DB.Table("auth").Create(&regEntity).Error; err != nil {
		ctx.JSON(handler.Throw500("Failed to create auth"))
		return
	}

	ctx.JSON(200, br.BaseSuccessResponse{
		Message: "success create auth",
		Success: true,
	})
}
