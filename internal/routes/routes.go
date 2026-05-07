package routes

import (
	"net/http"

	"go-advertiser-backend/internal/handlers"
	"go-advertiser-backend/internal/middlewares"
)

func RegisterRoutes(
	authHandler *handlers.AuthHandler,
) {
	http.HandleFunc("POST /api/register", authHandler.Register)
	http.HandleFunc("POST /api/login", authHandler.Login)
	http.HandleFunc(
		"GET /api/me",
		middlewares.AuthMiddleware(
			authHandler.Me,
		),
	)
}
