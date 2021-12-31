package models

type Todo struct {
	Id          string `json:id`
	CreatedAt   string `json:created_at`
	Deadline    string `json:deadline`
	Description string `json:description`
}

func NewTodo(id string, created string, deadline string, description string) *Todo {
	return &Todo{id, created, deadline, description}
}
