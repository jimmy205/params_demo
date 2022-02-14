package controllers

import (
	"params_demo/api/dtos"
	"params_demo/api/response"

	"github.com/gin-gonic/gin"
)

func ValidateUsername(c *gin.Context) {
	input := dtos.UsernameInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(err, input, c)
		return
	}

	response.SuccessWithMessage("validate username successs", c)
}

func ValidateAge(c *gin.Context) {
	input := dtos.AgeInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(err, input, c)
		return
	}

	response.SuccessWithMessage("validate age successs", c)
}
