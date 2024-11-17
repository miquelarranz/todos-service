package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/miquelarranz/todos-service/internal/handlers"
	"github.com/miquelarranz/todos-service/internal/persistence"
	"os"
)

func initDatabase() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	router := gin.Default()

	db := initDatabase()

	todoStore := persistence.NewTodoStore(db)
	todoHandler := handlers.NewTodoHandler(todoStore)

	router.GET("/todos", todoHandler.GetTodos)
	router.POST("/todos", todoHandler.PostTodos)
	router.GET("/todos/:id", todoHandler.GetTodo)

	router.Run(":8080")
}
