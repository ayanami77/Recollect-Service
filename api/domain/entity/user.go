package entity

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
)

type User struct {
	UserID    string    `json:"user_id" gorm:"primaryKey" validate:"required,min=3,max=20,alphanumunicode"`
	UserName  string    `json:"user_name" gorm:"not null"`
	Email     string    `json:"email" gorm:""`
	Password  string    `json:"password" gorm:"not null" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt time.Time `json:"deleted_at" gorm:""`
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
