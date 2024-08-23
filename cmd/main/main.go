package main

import (
	"github.com/sangharshseth/internal/handlers"
	"log"
	"net/http"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error Loading .env file")
// 	}
// }
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allow headers

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
	handlers.SetupVideoRoutes(mux)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", withCORS(mux)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
