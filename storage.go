package main

import (
	"context"
	"todo-api/models"

	"cloud.google.com/go/firestore"
)

func getFirebaseTodos(gfs *firestore.Client) []models.Todo {
	todos := make([]models.Todo, 0)
	docs := gfs.Collection("todos").Documents(context.Background())
	for doc, err := docs.Next(); err == nil; doc, err = docs.Next() {
		var todo models.Todo
		doc.DataTo(&todo)
		todos = append(todos, todo)
	}
	return todos
}

func addFirebaseTodos(gfs *firestore.Client, *todos models.Todo) {
	
}
