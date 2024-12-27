package users

import (
	// "gin-gorm/app/models"
	"gin-gorm/src/database"
	baseresponse "gin-gorm/src/utils/baseResponse"
	"gin-gorm/src/utils/handler"
	"gin-gorm/src/utils/logger"

	// "gin-gorm/utils/logger"
	"net/http"

	// "reflect"

	"github.com/gin-gonic/gin"
)

// type User struct {
// 	Username string `json:"username"`
// 	Email string `json:"email"`
// }

type Service struct {}
func (Service) GetAllUser(ctx *gin.Context) {
	var users []UserListResp 

	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message" : "error during select",
		})
		return
	}

	logger.Info("value", "p", users)
	ctx.JSON(http.StatusOK, baseresponse.BaseSuccessResponse[[]UserListResp]{
		Message: "success",
		Success: true,
		Data: users,
	})
}

func (Service) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	idUint, errId := handler.CheckID(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message" : "id must be a valid uint",
		})
		return
	}


	var user UserDetailResp
	err := database.DB.Table("users").Where("id = ?", idUint).Find(&user).Error

	logger.Debug("data", "user", err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "something went wrong!",
		})
		return
	}


	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message" : "no data found",
		})
		return
	}

	ctx.JSON(http.StatusOK, baseresponse.BaseSuccessResponse[UserDetailResp]{
		Message: "success mengambil detail",
		Success: true,
		Data: user,
	})

}

func (Service) CreateUser(ctx *gin.Context) {

	var userReq *UserCreateReq

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	// validate json
	handler.Validate(ctx, UserCreateDto, userReq)
	if ctx.IsAborted() {
		return
	}
	
	// validate email unique
	var emailExist UserDetailResp
	if err := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&emailExist).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	if emailExist.ID != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message" : "email already exists",
		})
		return
	}
	
	payload, errPayload := userReq.ToEntity()

	if errPayload != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message" : errPayload.Error(),
			"success" : false,
		})

		return 
	}

	if err := database.DB.Table("users").Create(&payload).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message" : err,
			"success" : false,
		})

		return 
	}

	// newBornDate, _ := time.Parse("2006-01-02", userReq.BornDate)
	// userReq.BornDate = time.str
	ctx.JSON(http.StatusOK, baseresponse.BaseSuccessResponse[any]{
		Message: "ok",
		Success: true,
		Data: userReq,
	})
}

func (Service) UpdateById (ctx *gin.Context) {
	// id := ctx.Param("id")

	// newID, errID := handler.CheckID(id)

	// if errID != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"message" : "id must be a valid uint",
	// 	})
	// 	return
	// }

}