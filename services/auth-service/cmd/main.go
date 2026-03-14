package main

import (
	"auth-service/internal/database"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	repo := &repository.UserRepository{DB: db}

	authService := &service.AuthService{Repo: repo}

	authHandler := &handler.AuthHandler{Service: authService}

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.Run(":8081")
}
