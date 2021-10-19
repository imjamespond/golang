package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"index"`
	Email        *string // pointer is nullable
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
