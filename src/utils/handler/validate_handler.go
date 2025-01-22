package handler

import (
	"fmt"
	// "gin-gorm/src/utils/logger"
	"net/http"
	// "reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	// "github.com/golodash/galidator/v2"
)

// func V(ctx *gin.Context, dto galidator.Validator, input any) {
// 	if errValidate := dto.Validate(ctx, input); errValidate != nil {

// 		errMsg := []string{}
// 		if vErr, ok := errValidate.(map[string]interface{}); ok {

// 			logger.Warn("error true validate")
// 			for _, item := range vErr {
// 				if k, okV := item.(map[string]interface{}); okV {
// 					// logger.Warn("naisuuu", "value", item)
// 					// errMsg = append(errMsg, k...)

// 					for _, i := range k {
// 						logger.Warn("naisuuu", "value", reflect.TypeOf(i))
// 						if j, okJ := i.([]string); okJ {
// 							errMsg = append(errMsg, j...)
// 						}
// 					}
// 				}
// 			}

// 		}
// 		// logger.Warn("error true validate", "data", errValidate)

// 		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
// 			"message": "Invalid validating",
// 			"success": false,
// 			"data":    errMsg,
// 		})
// 	}
// }

func translateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf("%s", e.Translate(trans)).Error()
		errs = append(errs, translatedErr)
	}
	return errs
}

func GetBody[T any](ctx *gin.Context) T {
	var b T

	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
			"success": false,
		})
	}
	
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	if err := v.Struct(b); err != nil {
		errors := translateError(err, trans)
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "invalid validation!",
			"success": false,
			"data":    errors,
		})
	}

	return b
}
