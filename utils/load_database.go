package utils

import (
	"github.com/nabinkatwal7/irlquest/db"
	"github.com/nabinkatwal7/irlquest/model"
)

func LoadDatabase(){
	db.Connect()

	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.Streak{})
}