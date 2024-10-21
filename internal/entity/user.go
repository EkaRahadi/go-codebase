package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId      uint64 `gorm:"primaryKey;column:user_id"`
	Username    string `gorm:"column:username;unique"`
	FullName    string `gorm:"column:full_name"`
	Email       string `gorm:"column:email;unique"`
	PhoneNumber string `gorm:"column:phone_number;unique"`
	Password    string `gorm:"column:password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
