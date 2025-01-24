package users

import (
	// "gin-gorm/app/models"
	"gin-gorm/src/database"
	"gin-gorm/src/database/dao"
	br "gin-gorm/src/utils/baseResponse"
	"gin-gorm/src/utils/handler"
	"gin-gorm/src/utils/logger"
	"time"

	// "gin-gorm/utils/logger"
	"net/http"

	// "reflect"

	"github.com/gin-gonic/gin"
)

// type User struct {
// 	Username string `json:"username"`
// 	Email string `json:"email"`
// }

type Service struct{}

// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string  "ok"
// @Router /string/{some_id} [get]
func (Service) GetAllUser(ctx *gin.Context) {

	var (
		users     []UserListResp
		totalUser int64
	)

	pg, errPagination := handler.GetPagination(ctx)
	if errPagination != nil {
		ctx.JSON(handler.Throw500(errPagination.Error()))
		return
	}

	err := database.DB.Table("users").Offset((pg.Page - 1) * pg.PageSize).Limit(pg.PageSize).Find(&users).Count(&totalUser).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error during select",
		})
		return
	}

	ctx.JSON(http.StatusOK, br.BaseSuccessResponsePagination{
		Message:    "success",
		Success:    true,
		Data:       users,
		Pagination: pg.GetResponse(int(totalUser)),
	})
}

func (Service) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	idUint, errId := handler.CheckID(id)
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be a valid uint",
		})
		return
	}

	var user UserDetailResp
	err := database.DB.Table("users").Where("id = ?", idUint).Find(&user).Error

	logger.Debug("data", "user", err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong!",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no data found",
		})
		return
	}

	ctx.JSON(http.StatusOK, br.BaseSuccessResponse{
		Message: "success mengambil detail",
		Success: true,
		Data:    user,
	})

}

func (Service) CreateUser(ctx *gin.Context) {

	// var userReq *UserCreateReq

	// if err := ctx.ShouldBindJSON(&userReq); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// validate json
	// handler.Validate(ctx, userReq)
	userReq := handler.GetBody[UserCreateReq](ctx)
	if ctx.IsAborted() {
		return
	}

	// validate email unique
	var emailExist UserDetailResp
	if err := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&emailExist).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if emailExist.ID != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "email already exists",
		})
		return
	}

	payload, errPayload := userReq.ToEntity()

	if errPayload != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": errPayload.Error(),
			"success": false,
		})

		return
	}

	if err := database.DB.Table("users").Create(&payload).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err,
			"success": false,
		})

		return
	}

	// newBornDate, _ := time.Parse("2006-01-02", userReq.BornDate)
	// userReq.BornDate = time.str
	ctx.JSON(http.StatusCreated, br.BaseSuccessResponse{
		Message: "ok",
		Success: true,
		Data:    userReq,
	})
}

func (Service) UpdateById(ctx *gin.Context) {
	var (
		id    uint
		errId error

		user dao.UserEntity
		// payload UserUpdateReq
	)

	if id, errId = handler.CheckID(ctx.Param("id")); errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id must a vaslid number",
		})
		return
	}

	if err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error; err != nil {
		ctx.JSON(handler.Throw422(err.Error()))
		return
	}

	if user.ID == nil {
		ctx.JSON(handler.Throw404("user not found"))
		return
	}

	// if err := ctx.ShouldBindJSON(&payload); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	logger.Error("input error")

	// 	return
	// }

	// handler.Validate(ctx, payload)
	payload := handler.GetBody[UserUpdateReq](ctx)
	if ctx.IsAborted() {
		return
	}

	if payload.Email != user.Email {
		var emailExist dao.UserEntity
		if err := database.DB.Table("users").Where("email = ?", payload.Email).Find(&emailExist).Error; err != nil {
			ctx.JSON(handler.Throw500("went wrong on getting email exist"))
			return
		}

		if emailExist.ID != nil {
			ctx.JSON(handler.Throw422("email is used"))
			return
		}
	}

	upPayload, errP := payload.ToEntity()
	if errP != nil {
		ctx.JSON(handler.Throw422(errP.Error()))
		return
	}

	upPayload.ID = &id
	upPayload.CreatedAt = user.CreatedAt
	upPayload.UpdatedAt = time.Now()

	if err := database.DB.Table("users").Save(&upPayload).Error; err != nil {
		ctx.JSON(handler.Throw422(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": upPayload,
	})
}

func (Service) DeleteById(ctx *gin.Context) {
	var (
		user dao.UserEntity
	)
	id, errId := handler.CheckID(ctx.Param("id"))

	if errId != nil {
		ctx.JSON(handler.Throw500("please enter valid id!"))
		return
	}

	if errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error; errFind != nil {
		ctx.JSON(handler.Throw500(errFind.Error()))
		return
	}

	if user.ID == nil {
		ctx.JSON(handler.Throw404("no user found"))
		return
	}

	if err := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&dao.UserEntity{}).Error; err != nil {
		ctx.JSON(handler.Throw500(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data berhasil di hapus",
	})
}
