package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ProcessImage(w http.ResponseWriter, r *http.Request){
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form-data
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a destination file where the video will be saved
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file content to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func SetupVideoRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /process-image", ProcessImage)
}