package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	email string
	password string
}

var db *sql.DB

func main() {
	var err error
	connStr := "user=postgres password=1234 dbname=db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/register", register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func isEmailValid(email string) {

}

func isPasswordSafe(password string) {
	
}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


}