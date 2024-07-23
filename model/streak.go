package model

import (
	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/irlquest/model"
	"gorm.io/gorm"
)

type Streak struct{
	gorm.Model
	UserID uint
	Days int `gorm:"default:0" json:"days"`
	Total int `gorm:"default:0" json:"total"`
}

func AddTodo(context *gin.Context){
	var input model.User

}