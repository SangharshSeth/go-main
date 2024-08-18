package handlers

import (
	"encoding/json"
	"net/http"

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
	defer file.Close()

	uploaded_image := storage.UploadImageToCloudinary(file)

	type Response struct {
		Status   string `json:"status"`
		ImageURL string `json:"image_url"`
	}

	response := Response{
		Status:   "success",
		ImageURL: uploaded_image,
	}

	sqsSender := queue.NewSqsSender("https://sqs.ap-south-1.amazonaws.com/873403696572/Image-Processor", "web-developer")
	sqsSender.SendMessageToSQS(uploaded_image)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func SetupVideoRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /process-image", ProcessImage)
}
