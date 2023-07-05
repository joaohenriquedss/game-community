package main

import (
	"api/usecase/game"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	userId := "1" // Value example, replace with desired user ID
	games := game.Filter(userId)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Initialization message of the API
		message := "API is up and running."
		w.Write([]byte(message))
	})

	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse, err := json.Marshal(games)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Print initialization message in the console
	log.Printf("API is up and running. URL: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
