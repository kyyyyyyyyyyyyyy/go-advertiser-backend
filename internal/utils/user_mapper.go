package utils

import (
	"go-advertiser-backend/internal/dto/response"
	"go-advertiser-backend/internal/models"
)

func ToUserResponse(user models.User) response.UserResponse {

	return response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
	}
}
