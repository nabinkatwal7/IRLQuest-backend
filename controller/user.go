package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/irlquest/db"
	"github.com/nabinkatwal7/irlquest/model"
)

func GetAllUsers(c *gin.Context){
	var users []model.User
	
	db.Database.Find(&users)
	c.JSON(200, users)
}