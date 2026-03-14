package routes

import (
	"io"
	"net/http"

	"gateway/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Public routes (no authentication)
	r.Any("/auth/*path", proxyAuth)

	// Protected routes
	todo := r.Group("/todos")
	todo.Use(middleware.JWTAuthMiddleware())
	{
		todo.Any("/*path", proxyTodo)
	}
}

func proxyAuth(c *gin.Context) {

	target := "http://auth-service:8081" + c.Param("path")

	proxyRequest(c, target)
}

func proxyTodo(c *gin.Context) {
	target := "http://todo-service:8082/todos" + c.Param("path")
	proxyRequest(c, target)
}

func proxyRequest(c *gin.Context, target string) {

	req, err := http.NewRequest(
		c.Request.Method,
		target,
		c.Request.Body,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Copy headers from original request
	req.Header = c.Request.Header

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", body)
}
