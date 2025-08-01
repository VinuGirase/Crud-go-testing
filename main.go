package main

import (
	"crud-test-go/config"
	"crud-test-go/routes"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config.InitDB() // ✅ Initialize DB

	r := routes.SetupRoutes()

	// CORS Setup
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(r)

	http.ListenAndServe(":8080", corsHandler)
}
