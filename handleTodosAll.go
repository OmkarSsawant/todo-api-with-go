package main

import (
	"encoding/json"
	"net/http"
	"todo-api/models"

	"github.com/gorilla/mux"
)

func HandleTodosAll(router *mux.Router, todos *[]models.Todo) {
	router.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}).Methods("GET")
}
