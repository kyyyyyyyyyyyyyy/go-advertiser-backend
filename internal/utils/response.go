package utils

import (
	"encoding/json"
	"net/http"

	"go-advertiser-backend/internal/dto/response"
)

func writeJSON(
	w http.ResponseWriter,
	statusCode int,
	payload response.APIResponse,
) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(
			w,
			`{"success":false,"message":"Internal Server Error"}`,
			http.StatusInternalServerError,
		)
	}
}

// ========================
// SUCCESS RESPONSE
// ========================

func Success(
	w http.ResponseWriter,
	statusCode int,
	message string,
	data interface{},
) {

	writeJSON(w, statusCode, response.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ========================
// ERROR RESPONSE
// ========================

func Error(
	w http.ResponseWriter,
	statusCode int,
	message string,
) {

	writeJSON(w, statusCode, response.APIResponse{
		Success: false,
		Message: message,
	})
}

// ========================
// VALIDATION ERROR
// ========================

func ValidationError(
	w http.ResponseWriter,
	errors map[string]string,
) {

	writeJSON(w, http.StatusBadRequest, response.APIResponse{
		Success: false,
		Message: "Validation error",
		Error:   errors,
	})
}

// ========================
// UNAUTHORIZED
// ========================

func Unauthorized(
	w http.ResponseWriter,
	message string,
) {

	writeJSON(w, http.StatusUnauthorized, response.APIResponse{
		Success: false,
		Message: message,
	})
}

// ========================
// NOT FOUND
// ========================

func NotFound(
	w http.ResponseWriter,
	message string,
) {

	writeJSON(w, http.StatusNotFound, response.APIResponse{
		Success: false,
		Message: message,
	})
}
