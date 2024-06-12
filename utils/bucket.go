package utils

import (
	"context"
	"fmt"

	"io"
	"log"
	"mime/multipart"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

func UploadToBucket(file *multipart.FileHeader, folder string) (string, error) {
	ctx := context.Background()
	
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucketName := "example-bucket-test-cc-trw" 
	destinationName := folder + "/" + uuid.NewString()
	imageName := filepath.Base(file.Filename)

	fileopen, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileopen.Close()

	// Read the file content into a byte slice
	fileBytes, err := io.ReadAll(fileopen)
	if err != nil {
		return "", err
	}

	wc := client.Bucket(bucketName).Object(destinationName).NewWriter(ctx)
	if _, err := wc.Write(fileBytes); err != nil {
		return "", fmt.Errorf("wc.Write: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("wc.Close: %v", err)
	}

	log.Printf("Image %s uploaded to %s/%s\n", imageName, bucketName, destinationName)
	return destinationName, nil
}
