package model

import (
	"time"
)

type User struct {
	ID         uint64     `gorm:"primaryKey"`
	Username   string     `gorm:"uniqueIndex;not null"`
	Passwd     string     `gorm:"not null"`
	CreateDate *time.Time `gorm:"index:,sort:desc"` // pass by pointer if it's posible to be null
}
