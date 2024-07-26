package model

import (
	"time"

	"encore.dev/types/uuid"
	"gorm.io/gorm"
)

type GoalStatus string 

const (
	Planned GoalStatus = "planned"
	InProgress GoalStatus = "in_progress"
	Completed GoalStatus = "completed"
)

type Goal struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID uint `json:"user_id"`
	Title string `gorm:"size:255;not null;" json:"title"`
	Description string `gorm:"size:255;" json:"description"`
	StartDate time.Time `json:"start_date"` 
	EndDate time.Time `json:"end_date"`
	Status GoalStatus `gorm:"type:enum('planned', 'in_progress', 'completed');default:'planned'" json:"status"`
}