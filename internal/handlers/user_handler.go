package handlers

import (
	"net/http"

	"go-advertiser-backend/internal/dto/response"
	"go-advertiser-backend/internal/services"
	"go-advertiser-backend/internal/utils"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUsers(
	w http.ResponseWriter,
	r *http.Request,
) {

	users, err := h.UserService.GetUsers()

	if err != nil {

		utils.Error(w, http.StatusInternalServerError, "Failed to fetch users")

		return
	}

	var userResponses []response.UserResponse

	for _, user := range users {
		userResponses = append(
			userResponses,
			utils.ToUserResponse(user),
		)
	}

	utils.Success(w, http.StatusOK, "Users fetched successfully", userResponses)
}
