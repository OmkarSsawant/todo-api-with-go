package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/models"

	"github.com/gorilla/mux"
)

func HandleTodosUpdate(router *mux.Router, todos *[]models.Todo) {
	router.HandleFunc("/api/todo/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		body := r.Body
		defer body.Close()
		var modTodo models.Todo
		err := json.NewDecoder(r.Body).Decode(&modTodo)
		if i := findIndex(*todos, id); i == -1 || err != nil {
			fmt.Fprint(w, "No Such Element In Data with id : ", id)
		} else {
			(*todos)[i] = modTodo
			jsonBytes, err := json.Marshal(todos)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			}
			saveJsonTodos(jsonBytes)
			fmt.Fprint(w, "Successfully Updated Todo (", id, ")")
		}
	}).Methods("PATCH")
}
