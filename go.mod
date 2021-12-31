module ingenuility.com/todo

go 1.17

replace todo-api/models => ./models

require (
	github.com/gorilla/mux v1.8.0
	todo-api/models v0.0.0-00010101000000-000000000000
)
