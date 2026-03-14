package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-service/internal/database"
	"todo-service/internal/handler"
	"todo-service/internal/repository"
	"todo-service/internal/routes"
	"todo-service/internal/service"
)

func main() {

	db, err := database.Connect()

	if err != nil {

		log.Fatal(err)
	}

	repo := &repository.TodoRepository{DB: db}

	svc := &service.TodoService{Repo: repo}

	h := &handler.TodoHandler{Service: svc}

	r := gin.Default()

	routes.RegisterTodoRoutes(r, h)

	r.Run(":8082")
}
