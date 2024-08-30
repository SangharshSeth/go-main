package webhooks

import (
	"encoding/json"
	"log"
	"net/http"
)

// WebhookPayload defines the structure of the payload sent to the webhook.
type WebhookPayload struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

// WebhookHandler processes incoming webhook requests.
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("%sWebhook received: %s%s", ColorGreen, payload.Data, ColorReset)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("Webhook received and processed successfully"))
	if err != nil {
		log.Printf("%sFailed to write response: %v%s", ColorRed, err, ColorReset)
	}
}
