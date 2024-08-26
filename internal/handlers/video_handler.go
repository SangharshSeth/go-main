package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/sangharshseth/internal/queue"
	"github.com/sangharshseth/internal/storage"
)

func ProcessImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Failed to close file: %v", cerr)
		}
	}()

	uploadedImage, err := storage.UploadImageToCloudinary(r.Context(), file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	type Response struct {
		Status   string `json:"status"`
		ImageURL string `json:"image_url"`
	}

	response := Response{
		Status:   "success",
		ImageURL: uploadedImage,
	}

	sqsSender := queue.NewSqsSender(os.Getenv("SQS_URL"), "web-developer")
	sqsSender.SendMessageToSQS(uploadedImage)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}

}

func SetupVideoRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /process-image", ProcessImage)
}
