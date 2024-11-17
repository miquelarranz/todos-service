package models

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Todos = []Todo{
	{ID: "1", Title: "Clean this", Description: "Clean the toilet"},
	{ID: "2", Title: "Go to the grocery store", Description: "I need milk and potatoes"},
	{ID: "3", Title: "Do some sport", Description: "Do the second session of the week"},
}

type Error struct {
	Message string `json:"message"`
}
