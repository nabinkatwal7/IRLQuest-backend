package model

import (
	"gorm.io/gorm"
)

type Streak struct{
	gorm.Model
	UserID uint
	Days int `gorm:"default:0" json:"days"`
	Total int `gorm:"default:0" json:"total"`
}

