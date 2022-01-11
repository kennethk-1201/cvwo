package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/kennethk-1201/cvwo/backend/internal/router"
)

func main() {
	r := router.Router()
	origins := handlers.AllowedOrigins([]string{"localhost:3000"})

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(origins)(r)))
}