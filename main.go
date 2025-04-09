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

	r.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var u user.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO users (email, age) VALUES (?, ?)", u.Email, u.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	}).Methods("POST")

	log.Println("Server started at :8080 port")
	log.Fatal(http.ListenAndServe(":8080", r))
}
