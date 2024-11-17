package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/miquelarranz/todos-service/internal/models"
	"github.com/miquelarranz/todos-service/internal/persistence"
)

type TodoHandler struct {
	todoStore *persistence.TodoStore
}

func NewTodoHandler(todoStore *persistence.TodoStore) *TodoHandler {
	return &TodoHandler{
		todoStore: todoStore,
	}
}

func (th *TodoHandler) GetTodos(c *gin.Context) {
	todos := models.Todos
	//todos, err := th.todoStore.FetchTodos(c)
	//if err != nil {
	//	c.IndentedJSON(http.StatusInternalServerError, models.Error{Message: "To-dos not found"})
	//}

	c.IndentedJSON(http.StatusOK, todos)
}

func (th *TodoHandler) PostTodos(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	models.Todos = append(models.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func (th *TodoHandler) GetTodo(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
