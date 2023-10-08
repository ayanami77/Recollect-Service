package response

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"time"
)

type UserResponse struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserResponse(user *entity.User) UserResponse {
	userResponse := UserResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse
}
