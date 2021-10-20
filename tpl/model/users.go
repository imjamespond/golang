package model

import (
	"time"
)

type Users1 struct {
	Id            int64      `gorm:"column:id"`
	Name          string     `gorm:"column:name"`
	Email         string     `gorm:"column:email"`
	Age           int16      `gorm:"column:age"`
	Birthday      *time.Time `gorm:"column:birthday"`
	Member_number string     `gorm:"column:member_number"`
	Activated_at  *time.Time `gorm:"column:activated_at"`
	Created_at    *time.Time `gorm:"column:created_at"`
	Updated_at    *time.Time `gorm:"column:updated_at"`
}

func (Users1) TableName() string {
	return "users"
}
