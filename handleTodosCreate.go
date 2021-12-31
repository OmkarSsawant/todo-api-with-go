package main

import (
	"encoding/json"
	"net/http"
	"todo-api/models"

	"github.com/gorilla/mux"
)

/*
	{
		"id" : "2",
		"deadline" : "31",
		"created_at" : "12",
		"description" : "Is it worth it to use GO"
	}
*/
func HandleTodosCreate(router *mux.Router, todos *[]models.Todo) {
	router.HandleFunc("/api/todo/create", func(w http.ResponseWriter, r *http.Request) {
		var body models.Todo
		json.NewDecoder(r.Body).Decode(&body)
		defer r.Body.Close()
		*todos = append(*todos, body)
		bytes, err := json.Marshal(todos)
		if err == nil {
			saveJsonTodos(bytes)
			json.NewEncoder(w).Encode(todos)
			return
		}
		json.NewEncoder(w).Encode(err)
	}).Methods("POST")
}
