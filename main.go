package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {

	sa := option.WithCredentialsFile("ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		log.Fatal(err)
	}
	gfs, err = app.Firestore(context.Background())
	todos := readJsonTodos("storage.json")

	router := mux.NewRouter()

	HandleTodosAll(router, &todos)
	HandleTodosCreate(router, &todos)
	HandleTodosDelete(router, &todos)
	HandleTodosUpdate(router, &todos)

	http.ListenAndServe(":8080", router)
}
