package main

import (
	"23-ci-cd/internal/db"
	"23-ci-cd/internal/user"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db := db.ConnMySQL()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO users (email, age) VALUES (?, ?)", user.Email, user.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
