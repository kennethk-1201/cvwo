package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
	_ "github.com/lib/pq"
	"github.com/kennethk-1201/cvwo/backend/internal/router"
)

func main() {
	r := router.Router()

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
    })

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", c.Handler(r)))
}