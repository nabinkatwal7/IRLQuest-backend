package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/irlquest/helper"
	"github.com/nabinkatwal7/irlquest/model"
)

func Login(context *gin.Context){
	var input model.AuthenticationInputLogin

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(*user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}