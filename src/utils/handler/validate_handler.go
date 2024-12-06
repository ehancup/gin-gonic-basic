package handler

import (
	"gin-gorm/src/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

func Validate(ctx *gin.Context, dto galidator.Validator, input any) {
	if errValidate := dto.Validate(ctx, input) ;errValidate != nil {

		errMsg := []string{}
		if vErr, ok := errValidate.(map[string]interface{});ok {

			logger.Warn("error true validate")
			for _, item := range vErr {
				if k, okV := item.([]string); okV {
					// logger.Warn("naisuuu", "value", item)
					errMsg = append(errMsg, k...)
				}
				// logger.Warn("naisuuu", "value", reflect.TypeOf(item))
			}

		}

		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message" : "Invalid validating",
			"success" : false,
			"data" : errMsg,
		})
	}
}

