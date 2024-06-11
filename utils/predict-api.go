package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func PredictionAPI(fileName string) (string, error) {
	// Define the request body (if needed)
	requestBody, err := json.Marshal(map[string]string{
		"imageName": fileName,
	})
	if err != nil {
		// fmt.Println("Error encoding JSON:", err)
		return err.Error(), err
	}

	// Define the URL of your Cloud Run service
	url := "http://127.0.0.1:8000/predict"
	// url := "https://py-load-h5-model-pc7eaxaqma-et.a.run.app/predict"

	// Make a POST request to the Cloud Run service
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		// fmt.Println("Error making POST request:", err)
		return err.Error(), err
	}
	defer response.Body.Close()

	// Read the response body
	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return err.Error(), err
	}

	// Print the prediction (or handle it as needed)
	floatValue := result["prediction"].(float64)
	return strconv.FormatFloat(floatValue, 'f', -1, 64), nil
}
