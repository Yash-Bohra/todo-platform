package handler

import (
	"net/http"
	"strconv"

	"todo-service/internal/models"
	"todo-service/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Service *service.TodoService
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {

	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err := h.Service.CreateTodo(&todo)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {

	todos, err := h.Service.GetTodos()

	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, todos)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {

	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	todo, err := h.Service.GetTodo(id)

	if err != nil {

		c.JSON(404, gin.H{"error": "todo not found"})

		return
	}

	c.JSON(200, todo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {

	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var todo models.Todo

	c.BindJSON(&todo)

	todo.ID = id

	err := h.Service.UpdateTodo(&todo)

	if err != nil {

		c.JSON(500, gin.H{"error": "update failed"})

		return
	}

	c.JSON(200, gin.H{"message": "updated"})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {

	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	err := h.Service.DeleteTodo(id)

	if err != nil {

		c.JSON(500, gin.H{"error": "delete failed"})

		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}
