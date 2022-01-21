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
        AllowedOrigins: []string{"xd","http://localhost:3000", "http://192.168.43.158", "http://192.168.43.158:3000", "https://cvwo-kenneth.netlify.app"},
        AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    })

	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", c.Handler(r)))
}