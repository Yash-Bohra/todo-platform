package routes

import (
	"todo-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine, h *handler.TodoHandler) {

	todos := r.Group("/todos")

	todos.POST("/", h.CreateTodo)

	todos.GET("/", h.GetTodos)

	todos.GET("/:id", h.GetTodo)

	todos.PUT("/:id", h.UpdateTodo)

	todos.DELETE("/:id", h.DeleteTodo)
}
