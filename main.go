package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/models"

	"github.com/gorilla/mux"
)

func main() {
	todos := readJsonTodos("storage.json")

	router := mux.NewRouter()
	HandleTodosAll(router, &todos)
	HandleTodosCreate(router, &todos)
	HandleTodosDelete(router, &todos)

	router.HandleFunc("/api/todo/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		body := r.Body
		defer body.Close()
		var modTodo models.Todo
		err := json.NewDecoder(r.Body).Decode(modTodo)
		if i := findIndex(todos, id); i == -1 || err != nil {
			fmt.Fprint(w, "No Such Element In Data with id : ", id)
		} else {
			todos[i] = modTodo
			jsonBytes, err := json.Marshal(todos)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			}
			saveJsonTodos(jsonBytes)
			fmt.Fprint(w, "Successfully Updated Todo (", id, ")")
		}
	}).Methods("PATCH")

	http.ListenAndServe(":8080", router)
}

func findIndex(p []models.Todo, id string) int {
	for i, todo := range p {
		if todo.Id == id {
			return i
		}
	}
	return -1
}

func remove(p []models.Todo, id string) (bool, []models.Todo) {
	i := findIndex(p, id)
	if i == -1 {
		return false, p
	}
	modList := p[:i]
	modList = append(modList, p[i+1:]...)
	return true, modList
}
