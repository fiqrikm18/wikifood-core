package models

import (
	"gorm.io/gorm"
	"time"
)

type token struct {
	Id                  uint `gorm:"primaryKey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	AccesToken          string
	RefreshToken        string
	AccessTokenExpired  time.Time
	RefreshTokenExpired time.Time
}
