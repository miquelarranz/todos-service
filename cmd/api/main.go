package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miquelarranz/todos-service/internal/todo"
)

func main() {
	router := gin.Default()

	router.GET("/todos", todo.GetTodos)
	router.POST("/todos", todo.PostTodos)
	router.GET("/todos/:id", todo.GetTodo)

	router.Run(":8080")
}
