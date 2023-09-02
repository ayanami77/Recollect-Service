package entity

import (
	"time"
)

type User struct {
	UserID    string    `json:"user_id" gorm:"primaryKey"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt time.Time `json:"deleted_at"`
}
