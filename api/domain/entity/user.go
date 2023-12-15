package entity

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
)

type User struct {
	Sub                         string    `json:"sub" gorm:"primaryKey"` //Google OAuthの識別子
	UserID                      string    `json:"user_id" gorm:"unique;not null" validate:"required,min=3,max=20,alphanumunicode"`
	UserName                    string    `json:"user_name" gorm:"not null"`
	Email                       string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	ComprehensiveAnalysisResult string    `json:"comprehensive_analysis_result"`
	ComprehensiveAnalysisScore  string    `json:"comprehensive_analysis_score"`
	CreatedAt                   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt                   time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt                   time.Time `json:"deleted_at" gorm:""`
	Cards                       []Card    `json:"cards" gorm:"foreignKey:Sub;references:Sub"`
}

// TODO
func IncludeNumeric(fl validator.FieldLevel) bool {
	return checkRegexp("[0-9]", fl.Field().String())
}

func IncludeAlphabetic(fl validator.FieldLevel) bool {
	return checkRegexp("^[a-zA-Z]+$", fl.Field().String())
}

func checkRegexp(reg, str string) bool {
	r := regexp.MustCompile(reg).Match([]byte(str))
	return r
}

//https://qiita.com/syoimin/items/b3923fea6070b0a3df8f
