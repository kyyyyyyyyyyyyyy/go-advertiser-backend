package services

import (
	"go-advertiser-backend/internal/models"
	"go-advertiser-backend/internal/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepo.FindAll()
}
