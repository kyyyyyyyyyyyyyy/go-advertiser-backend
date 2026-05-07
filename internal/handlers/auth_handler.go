package handlers

import (
	"encoding/json"
	"net/http"

	"go-advertiser-backend/internal/dto/request"
	"go-advertiser-backend/internal/dto/response"
	"go-advertiser-backend/internal/middlewares"
	"go-advertiser-backend/internal/services"
	"go-advertiser-backend/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(
	authService *services.AuthService,
) *AuthHandler {

	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req request.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		utils.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)

		return
	}

	// validate request
	if err := validate.Struct(req); err != nil {

		validationErrors := utils.FormatValidationErrors(err)

		utils.ValidationError(
			w,
			validationErrors,
		)

		return
	}

	user, err := h.AuthService.Register(req)
	if err != nil {

		utils.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	utils.Success(
		w,
		http.StatusCreated,
		"register success",
		utils.ToUserResponse(*user),
	)
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		utils.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)

		return
	}

	// validate request
	if err := validate.Struct(req); err != nil {

		utils.ValidationError(
			w,
			utils.FormatValidationErrors(err),
		)

		return
	}

	token, err := h.AuthService.Login(req)
	if err != nil {

		utils.Error(
			w,
			http.StatusUnauthorized,
			err.Error(),
		)

		return
	}

	authResponse := response.AuthResponse{
		AccessToken: token,
	}

	utils.Success(
		w,
		http.StatusOK,
		"login success",
		authResponse,
	)
}

func (h *AuthHandler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value(
		middlewares.UserIDKey,
	).(uuid.UUID)

	if !ok {

		utils.Unauthorized(
			w,
			"Unauthorized",
		)

		return
	}

	user, err := h.AuthService.Me(userID)
	if err != nil {

		utils.Error(
			w,
			http.StatusNotFound,
			err.Error(),
		)

		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Current user fetched successfully",
		utils.ToUserResponse(*user),
	)
}
