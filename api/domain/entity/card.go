package entity

import "time"

type Card struct {
	CardID    string    `json:"card_id" gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"not null;index"`
	Period    string    `json:"period" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:""`
	Tags      string    `json:"tags" gorm:""` // TODO: []stringだとエラーになる
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt time.Time `json:"deleted_at" gorm:""`
	User      User      `gorm:"foreignKey:UserID"`
}
