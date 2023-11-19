package entity

import (
	"github.com/lib/pq"
	"time"
)

type Card struct {
	CardID         string         `json:"card_id" gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	Sub            string         `json:"sub" gorm:"not null;index"`
	Period         string         `json:"period" gorm:"not null"`
	Title          string         `json:"title" gorm:"not null"`
	Content        string         `json:"content" gorm:""`
	AnalysisResult string         `json:"analysis_result" gorm:""`
	Tags           pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt      time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt      time.Time      `json:"deleted_at" gorm:""`
	User           User           `gorm:"references:Sub;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
