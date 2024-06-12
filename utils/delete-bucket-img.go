package utils

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func DeleteFromBucket(folderPath string, fileName string) error {
	ctx := context.Background()
	serviceAccountKeyFile := "../bangkit-cloud-computing-2af7d72444a8.json"
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountKeyFile))
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucketName := "example-bucket-test-cc-trw"

	// Create the full path to the object
	objectPath := folderPath + "/" + fileName
	fmt.Print("\n" + objectPath + "\n")

	// Create a handle to the object in the bucket
	objectHandle := client.Bucket(bucketName).Object(objectPath)

	// Attempt to delete the object
	if err := objectHandle.Delete(ctx); err != nil {
		return fmt.Errorf("objectHandle.Delete: %v", err)
	}

	log.Printf("File %s deleted from %s/%s\n", fileName, bucketName, objectPath)
	return nil
}
