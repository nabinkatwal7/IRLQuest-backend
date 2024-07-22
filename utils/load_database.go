package utils

import "github.com/nabinkatwal7/irlquest/db"

func LoadDatabase(){
	db.Connect()
}