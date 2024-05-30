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
	"google.golang.org/api/option"
)

func UploadToBucket(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	serviceAccountKeyFile := "../bangkit-cloud-computing-2af7d72444a8.json"
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountKeyFile))
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucketName := "example-bucket-test-cc-trw"
	destinationName := uuid.NewString()
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
