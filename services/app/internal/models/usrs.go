package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	IsOnline bool `gorm:"default:false"`
	IsBanned bool `gorm:"default:false"`

	HP int `gorm:"default:120"`

	Position Position `gorm:"foreignKey:UserID"`

	LastSeen time.Time
}