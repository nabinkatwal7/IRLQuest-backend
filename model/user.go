package model

import (
	"github.com/google/uuid"
	"github.com/nabinkatwal7/irlquest/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ID string `gorm:"primaryKey"`
	FirstName string `gorm:"size:255;not null;" json:"firstName"`
	LastName string `gorm:"size:255;not null;" json:"lastName"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255; not null;" json:"-"`
	Avatar string `gorm:"size:255;" json:"avatar"`
	Theme string `gorm:"size:255;" json:"theme"`
	Streaks []Streak `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"streaks"`
	Goals []Goal `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"goals"`
}


func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
  // UUID version 4
  user.ID = uuid.NewString()
  return
}

func (user *User) Save() (*User, error){
	err := db.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	return nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	err := db.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, err
}
 
func FindUserById(id uint) (*User, error) {
	var user User
	err := db.Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, err
}

func FindUserByUsername (username string) (*User, error) {
	var user User
	err := db.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &User{}, err	
	}
	return &user, err
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}