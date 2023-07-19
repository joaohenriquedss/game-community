package main

import (
	db "api/adapter/sql/config"
	adapterUser "api/adapter/user"
	"api/usecase/game"
	usecaseUser "api/usecase/user"
	usersDtos "api/usecase/user/dtos"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func initUsersMock(usecase *usecaseUser.ProcessUser) {
	for i := 0; i < 100; i++ {
		input := usersDtos.UserCreateDto{
			Name:            fmt.Sprintf("User %d", i),
			Email:           fmt.Sprintf("email%d@gmail.com", i),
			Password:        "12345678",
			PasswordConfirm: "12345678",
		}
		_, err := usecase.Execute(input)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {
	userId := "1"
	games := game.Filter(userId)
	db, err := db.ConnectToDB()
	userRepo := adapterUser.NewUserRepositoryDB(db)
	usecase := usecaseUser.NewProcessUser(userRepo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := "API is up and running."
		w.Write([]byte(message))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := usecase.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResponse, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
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
