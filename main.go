package main

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"log"
	"net/http"
)

func main() {
	app, err := firebase.NewApp(context.Background(), nil) // firebase will automatically read auth token from environment variable GOOGLE_APPLICATION_CREDENTIALS
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	http.Handle("/notify", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			var data map[string]string
			err := json.NewDecoder(request.Body).Decode(&data)
			if err != nil {
				log.Printf("error decoding request body: %v\n", err)
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			// create a message to send to the device
			_, err = client.Send(context.Background(), &messaging.Message{
				Notification: &messaging.Notification{
					Title: "New order",
				},
				Data: map[string]string{
					"test": "test",
				},
				Token: "fcm-token-here"})
			if err == nil {
				writer.WriteHeader(http.StatusOK)
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}

		}
	}))

	http.ListenAndServe(":8080", nil)
}
