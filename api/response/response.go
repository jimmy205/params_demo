package response

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BadRequest(err error, input interface{}, ctx *gin.Context) {

	validErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	var missingParams []string
	for _, e := range validErrs {
		field, ok := reflect.TypeOf(input).FieldByName(e.Field())
		if !ok {
			missingParams = append(missingParams, "undefined")
			continue
		}
		missingParams = append(missingParams, field.Tag.Get("json"))
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("[ 參數錯誤 ] %s", strings.Join(missingParams, ",")),
	})
	return
}

func InternalError(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func SuccessWithData(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func SuccessWithMessage(message string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
