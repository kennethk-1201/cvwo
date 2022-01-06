package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "encoding/json"
// "math/rand"
// "strconv"

func main() {
	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/task", getTasks).Methods("GET")
	r.HandleFunc("/api/task/{id}", getTask).Methods("GET")
	r.HandleFunc("/api/task/add", createTask).Methods("POST")
	r.HandleFunc("/api/task/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/api/task/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getTasks() {

}

func getTask() {

}

func createTask() {

}

func updateTask() {

}

func deleteTask() {

}
