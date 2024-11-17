package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/miquelarranz/todos-service/internal/handlers"
	"github.com/miquelarranz/todos-service/internal/persistence"

	"os"
)

func main() {
	router := gin.Default()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err == nil {
		panic(err)
	}

	todoStore := persistence.NewTodoStore(db)
	todoHandler := handlers.NewTodoHandler(todoStore)

	router.GET("/todos", todoHandler.GetTodos)
	router.POST("/todos", todoHandler.PostTodos)
	router.GET("/todos/:id", todoHandler.GetTodo)

	router.Run(":8080")
}
