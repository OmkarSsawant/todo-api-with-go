package main

import "todo-api/models"

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
