package services

import (
	"errors"
	"strings"

	"go-advertiser-backend/internal/dto/request"
	"go-advertiser-backend/internal/models"
	"go-advertiser-backend/internal/repositories"
	"go-advertiser-backend/internal/utils"

	"github.com/google/uuid"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(
	userRepo *repositories.UserRepository,
) *AuthService {

	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(
	req request.RegisterRequest,
) (*models.User, error) {
	// sanitize
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Name = strings.TrimSpace(req.Name)

	// check existing email
	existingUser, err := s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// create user
	user, err := s.UserRepo.Create(
		req.Name,
		req.Email,
		hashedPassword,
		req.Role,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(
	req request.LoginRequest,
) (string, error) {

	req.Email = strings.TrimSpace(
		strings.ToLower(req.Email),
	)

	// find user
	user, err := s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("invalid credentials")
	}

	// verify password
	err = utils.VerifyPassword(
		req.Password,
		user.Password,
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// generate jwt
	token, err := utils.GenerateJWT(
		user.ID,
		string(user.Role),
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Me(
	userID uuid.UUID,
) (*models.User, error) {

	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
