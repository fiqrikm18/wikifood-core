package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string
	Email     string
	Password  string
}
