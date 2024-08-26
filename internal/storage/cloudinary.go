package storage

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

// Global Cloudinary client
var cld *cloudinary.Cloudinary

// Initialize the Cloudinary client once
func init() {
	var err error
	cld, err = cloudinary.New()
	if err != nil {
		// Handle initialization error, perhaps panic or log it
		panic(err)
	}
}

func UploadImageToCloudinary(ctx context.Context, file multipart.File) (string, error) {
	uploadResult, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{
			PublicID:       uuid.New().String(),
			UseFilename:    api.Bool(true),
			UniqueFilename: api.Bool(true),
			Overwrite:      api.Bool(false),
		})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
