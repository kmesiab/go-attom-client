package nationalflooddata

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestFloodDataUnmarshal(t *testing.T) {

	// Load the JSON file
	filePath := "../../samples/nationalflooddata/data.json"
	file, err := os.Open(filePath)

	if err != nil {
		t.Fatalf("Failed to open JSON file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatalf("Failed to close JSON file: %v", err)
		}
	}(file)

	// Decode the JSON data into the FloodData struct
	var floodData FloodData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&floodData); err != nil {
		t.Fatalf("Failed to decode JSON data: %v", err)
	}

	// Perform basic checks to ensure data is unmarshaled correctly
	if floodData.Geocode.Relevance != 1 {
		t.Errorf("Expected Geocode.Relevance to be 1, got %d", floodData.Geocode.Relevance)
	}

	if floodData.Result.Loma[0].CaseNumber != "14-05-8454A" {
		t.Errorf("Expected Loma[0].CaseNumber to be '14-05-8454A', got %s", floodData.Result.Loma[0].CaseNumber)
	}

	if floodData.Result.Property.Sqft != "1200" {
		t.Errorf("Expected Property.Sqft to be '1200', got %s", floodData.Result.Property.Sqft)
	}

	output, err := json.MarshalIndent(floodData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal FloodData struct: %v", err)
	}

	fmt.Println(string(output))

}
