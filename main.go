package main

import (
	"context"
	"log"
	"net/http"
)

func main() {

	app, err := firebase.NewApp(context.Background(), nil) // firebase will automatically read auth token from environment variable GOOGLE_APPLICATION_CREDENTIALS
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	http.Handle("/notify", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":

		}
	}))

	http.ListenAndServe(":8080", nil)
}
