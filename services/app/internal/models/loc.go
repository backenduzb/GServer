package models

import (
	"time"
)

type Position struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"UniqueIndex"`

	X float64 `gorm:"default: 0"`
	Y float64 `gorm:"default: 0"`
	Z float64 `gorm:"default: 0"`

	UpdatedAt time.Time
}