package main

import (
	"encoding/json"
	"os"
	"todo-api/models"
)

func saveJsonTodos(data []byte) bool {
	err := os.WriteFile("storage.json", data, os.ModeTemporary)
	if err != nil {
		return false
	}
	return err == nil
}

func readJsonTodos(fileName string) []models.Todo {
	var todos []models.Todo
	jsonBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}
	json.Unmarshal(jsonBytes, todos)
	return todos
}
