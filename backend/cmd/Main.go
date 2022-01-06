package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kennethk-1201/cvwo/backend/internal/helper"
)

// "encoding/json"
// "math/rand"
// "strconv"

// Task Struct (Model)
type Task struct {
	ID      string `json:"id"`
	Isbn    string `json:"isbn"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Set Environment Variables
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	fmt.Println(helper.GetEnvVariable("TEST_VAR"))
	// Route Handlers / Endpoints
	r.HandleFunc("/read", getTasks).Methods("GET")
	r.HandleFunc("/read/{id}", getTask).Methods("GET")
	r.HandleFunc("/create", createTask).Methods("POST")
	r.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {

}

// get specific task
func getTask(w http.ResponseWriter, r *http.Request) {

}

// create task
func createTask(w http.ResponseWriter, r *http.Request) {

}

// update task
func updateTask(w http.ResponseWriter, r *http.Request) {

}

// delete task
func deleteTask(w http.ResponseWriter, r *http.Request) {

}
