package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/sangharshseth/internal/handlers"
	"github.com/sangharshseth/internal/webhooks"
)

func init() {
	fmt.Print("THIS EXECUTES FIRST")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, r.Proto, time.Since(start))
	})
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	// Set up routes
	handlers.ImageProcessor(mux)

	mux.HandleFunc("/webhook", webhooks.WebhookHandler)
	mux.HandleFunc("/websocket", handlers.WebSockets)

	handler := loggerMiddleware(withCORS(mux))

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
