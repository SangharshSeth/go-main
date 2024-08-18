package storage

import (
	"context"
	"log"
	"mime/multipart"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go/v2/api"
)

func UploadImageToCloudinary(file multipart.File) string{
	cld, err := cloudinary.New()
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary, %v", err)
	}
	var ctx = context.Background()
	uploadResult, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{
			PublicID:       "image2",
			UseFilename:    *api.Bool(true),
			UniqueFilename: *api.Bool(true),
			Overwrite:      *api.Bool(true)})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	return uploadResult.SecureURL
}