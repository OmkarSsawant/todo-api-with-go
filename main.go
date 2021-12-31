package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	todos := readJsonTodos("storage.json")

	router := mux.NewRouter()

	HandleTodosAll(router, &todos)
	HandleTodosCreate(router, &todos)
	HandleTodosDelete(router, &todos)
	HandleTodosUpdate(router, &todos)

	http.ListenAndServe(":8080", router)
}
