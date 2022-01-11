package router

import (

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/kennethk-1201/cvwo/backend/internal/middleware"
)

func Router() *mux.Router{
	// Init Router
	r := mux.NewRouter()

	// Set Environment Variables
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	// Route middleware / Endpoints
	r.HandleFunc("/read", middleware.GetTasks).Methods("GET")
	r.HandleFunc("/read/{id}", middleware.GetTask).Methods("GET")
	r.HandleFunc("/create", middleware.CreateTask).Methods("POST")
	r.HandleFunc("/update/{id}", middleware.UpdateTask).Methods("PUT")
	r.HandleFunc("/delete/{id}", middleware.DeleteTask).Methods("DELETE")

	return r
}