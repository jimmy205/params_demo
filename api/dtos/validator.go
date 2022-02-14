package dtos

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	for _, f := range validatorFuncs {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			if err := v.RegisterValidation(functionName(f), f); err != nil {
				panic(fmt.Sprintf("init validation func error:%s", err))
			}
		}
	}
}

func functionName(f interface{}) string {
	path := strings.Split(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".",
	)
	return path[len(path)-1]
}

var validatorFuncs = []func(validator.FieldLevel) bool{
	username, password, timing,
}

func username(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	matchbool, _ := regexp.MatchString(`[a-zA-Z0-9]{6,10}$`, name)
	return matchbool
}

func password(fl validator.FieldLevel) bool {
	return true
}

func timing(fl validator.FieldLevel) bool {
	return true
}
