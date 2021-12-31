package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/models"

	"github.com/gorilla/mux"
)

func HandleTodosDelete(router *mux.Router, todos *[]models.Todo) {
	router.HandleFunc("/api/todo/del/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		fmt.Printf("seeking : %s", id)
		didUpdate, modTodos := remove(*todos, id)
		*todos = modTodos
		jsonBytes, err := json.Marshal(modTodos)
		if err != nil && didUpdate {
			fmt.Fprint(w, "Operation not succeded ", id)
		} else {
			saveJsonTodos(jsonBytes)
			fmt.Fprint(w, "Success fully deleted ", id)
		}
	}).Methods("DELETE")
}
