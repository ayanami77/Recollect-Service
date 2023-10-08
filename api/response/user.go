package response

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"time"
)

type UserResponse struct {
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
