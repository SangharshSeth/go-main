package handlers

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/sangharshseth/internal/queue"
	"github.com/sangharshseth/internal/storage"
)

// ImageOptions holds the options for image processing.
type ImageOptions struct {
	CropWidth  string `json:"cropWidth,omitempty"`
	CropHeight string `json:"cropHeight,omitempty"`
	Type       string `json:"option,omitempty"`
}

// QueueMessageData represents the data sent to the SQS queue.
type QueueMessageData struct {
	ImageOptions *ImageOptions `json:"imageOptions,omitempty"`
	ImageURL     string        `json:"imageURL"`
}

// ProcessImage handles the image processing request.
func ProcessImage(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form with a maximum memory of 10 MB.
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the uploaded image file.
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer closeFile(file)

	// Process the options from the form data.
	options := processOptions(r)

	// Upload the image to Cloudinary.
	uploadedImage, err := storage.UploadImageToCloudinary(r.Context(), file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare data for SQS.
	queueData := &QueueMessageData{
		ImageOptions: options,
		ImageURL:     uploadedImage,
	}

	// Send the image URL and options to SQS.
	if err := sendToQueue(queueData); err != nil {
		log.Printf("Failed to send message to SQS: %v", err)
		http.Error(w, "Failed to process the request", http.StatusInternalServerError)
		return
	}

	// Create and send the JSON response.
	sendJSONResponse(w, uploadedImage, options)
}

// closeFile closes the uploaded file and logs an error if it fails.
func closeFile(file multipart.File) {
	if err := file.Close(); err != nil {
		log.Printf("Failed to close file: %v", err)
	}
}

// processOptions retrieves and processes options from the form data.
func processOptions(r *http.Request) *ImageOptions {
	options := &ImageOptions{}

	switch r.FormValue("option") {
	case "crop":
		options.CropWidth = r.FormValue("cropWidth")
		options.CropHeight = r.FormValue("cropHeight")
		options.Type = "crop"
	default:
		log.Printf("Unexpected option: %s", r.FormValue("option"))
	}

	log.Printf("Processed options: %+v", options)
	return options
}

// sendToQueue sends the image URL and options as JSON to the SQS queue.
func sendToQueue(queueData *QueueMessageData) error {
	sqsSender := queue.NewSqsSender(os.Getenv("SQS_URL"), "web-developer")

	// Create a message payload including the image URL and options.
	messagePayload, err := json.Marshal(queueData)
	if err != nil {
		return err
	}

	// Send the JSON message to SQS.
	log.Printf("Message sent is: %s\n", string(messagePayload))
	sqsSender.SendMessageToSQS(string(messagePayload))
	return nil
}

// sendJSONResponse sends a JSON response back to the client.
func sendJSONResponse(w http.ResponseWriter, imageURL string, options *ImageOptions) {
	response := struct {
		Status   string        `json:"status"`
		ImageURL string        `json:"image_url"`
		Options  *ImageOptions `json:"options,omitempty"`
	}{
		Status:   "success",
		ImageURL: imageURL,
		Options:  options,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SetupVideoRoutes sets up the routes for video processing.
func SetupVideoRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/process-image", ProcessImage)
}
