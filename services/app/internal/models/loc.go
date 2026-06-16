package models

import (
	"time"
)

type Position struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"UniqueIndex"`

	X float64 
	Y float64
	Z float64 

	UpdatedAt time.Time
}