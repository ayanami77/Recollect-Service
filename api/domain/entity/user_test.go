package entity

import (
	"github.com/go-playground/validator/v10"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	sampleTime := time.Date(2023, 9, 1, 0, 0, 0, 0, time.Local)
	okTests := []User{
		{UserID: "aaa000", UserName: "aaa000", Email: "", Password: "aaa000", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "aaa", UserName: "", Email: "", Password: "aaaaaa", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "xxxxxxxxxxxxxxxxxxxx", UserName: "zzzzzz", Email: "", Password: "aaaaaaaaaaaaaa", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
	}

	ngTests := []User{
		{UserID: "", UserName: "", Email: "", Password: "aaa000", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "aa", UserName: "", Email: "", Password: "aaa000", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "aaaaaaaaaaaaaaaaaaaaa", UserName: "", Email: "", Password: "aaa000", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "aaa%", UserName: "", Email: "", Password: "aaa000", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "aaa", UserName: "", Email: "", Password: "", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
		{UserID: "aaa", UserName: "", Email: "", Password: "aaa00", CreatedAt: sampleTime, UpdatedAt: sampleTime, DeletedAt: sampleTime},
	}

	validate := validator.New()
	//validate.RegisterValidation("includeNumeric", IncludeAlphabetic)
	//validate.RegisterValidation("includeAlphabetic", IncludeNumeric)

	for _, tc := range okTests {
		err := validate.Struct(tc)
		if err != nil {
			t.Error(err)
		}
	}

	for _, tc := range ngTests {
		err := validate.Struct(tc)
		if err == nil {
			t.Error(tc)
		}
	}
}
