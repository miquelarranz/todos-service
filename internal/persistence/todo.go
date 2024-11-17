package persistence

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/miquelarranz/todos-service/internal/models"
	"log"
)

type TodoStore struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{
		db: db,
	}
}

func (ts *TodoStore) FetchTodos(c *gin.Context) ([]models.Todo, error) {
	rows, err := ts.db.QueryContext(c, "SELECT * FROM todos")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	todos := make([]models.Todo, 0)

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return todos, nil
}
