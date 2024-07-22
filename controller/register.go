package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/irlquest/model"
)

func Register(c *gin.Context) {
	var input model.AuthenticationInputRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Username: input.Username,
		Email: input.Email,
		Password: input.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON((http.StatusCreated), savedUser)
}