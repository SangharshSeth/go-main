package main

import (
	"github.com/joho/godotenv"
	"github.com/sangharshseth/internal/database"
	"github.com/sangharshseth/internal/handlers"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func main() {
	os.Getenv()
	database.ConnectMongo(os.Getenv("MONGO_URI"))
	defer database.DisconnectMongo()
	mux := http.NewServeMux()
	// Set up routes
	handlers.SetupVideoRoutes(mux)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
