package main

import (
	"log"
	"net/http"
	"os"

	"go-advertiser-backend/internal/config"
	"go-advertiser-backend/internal/handlers"
	"go-advertiser-backend/internal/repositories"
	"go-advertiser-backend/internal/routes"
	"go-advertiser-backend/internal/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed load .env")
	}

	port := os.Getenv("APP_PORT")

	db := config.ConnectDB()

	// repositories
	userRepo := repositories.NewUserRepository(db)

	// services
	authService := services.NewAuthService(userRepo)

	// handlers
	authHandler := handlers.NewAuthHandler(authService)

	// routes
	routes.RegisterRoutes(authHandler)

	log.Println("🚀 Server running on :" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
