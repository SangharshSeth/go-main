package main

import (
	"github.com/joho/godotenv"
	"github.com/sangharshseth/internal/handlers"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func main() {

	mux := http.NewServeMux()
	// Set up routes
	handlers.SetupRoutes(mux)
	handlers.SetupVideoRoutes(mux)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
