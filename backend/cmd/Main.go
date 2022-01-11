package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/kennethk-1201/cvwo/backend/internal/helper"
)

// "encoding/json"
// "math/rand"
// "strconv"

// Task Struct (Model)
type Task struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Body string `json:"body"`
	Deadline string `json:"deadline"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Task `json:"data"`
    Message string `json:"message"`
}

func checkErr(err error) {
	if err != nil {
        log.Fatalf("Error opening database: %q", err)
	}
} 


func setupDB() *sql.DB {
	os.Setenv("DATABASE_URL", "postgres://rbsgtjevpzvhyr:dfde88c49176d084a4c0000cb74d0c2f762f09806e39bafea80d26d5b5032335@ec2-34-249-49-9.eu-west-1.compute.amazonaws.com:5432/d8hae478lmhvj0")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	checkErr(err)

    return db
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
	db := setupDB()

	printMessage("Getting tasks...")

	rows, err := db.Query("SELECT * FROM movies")
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
