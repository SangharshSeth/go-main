package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2/api"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

// Global Cloudinary client
// var cld *cloudinary.Cloudinary

func init() {
	fmt.Print(os.Getenv("CLOUDINARY_URL"))
	// var err error
	// cld, err = cloudinary.New()
	// if err != nil {
	// 	// Handle initialization error, perhaps panic or log it
	// 	panic(err)
	// }
}

func UploadImageToCloudinary(ctx context.Context, file multipart.File) (string, error) {
	cld, err := cloudinary.New()
	if err != nil {
		// Handle initialization error, perhaps panic or log it
		panic(err)
	}
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
