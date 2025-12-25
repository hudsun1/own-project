package main

import (
	"fmt"
	"log"

	"BubbleTesting/api"
	"BubbleTesting/config"
	"BubbleTesting/models"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create API client
	client := api.NewClient(cfg)

	// Example: Create a record
	data := map[string]interface{}{
		"name":        "Test Record",
		"description": "This is a test record",
	}

	resp, err := client.CreateRecord("your_object_type", data)
	if err != nil {
		log.Fatalf("Error creating record: %v", err)
	}
	fmt.Printf("Created record: %+v\n", resp)

	// Example: Query records
	query := models.NewBubbleQuery()
	query.AddConstraint("name", "equals", "Test Record")
	query.AddSort("Created Date", "desc")
	query.Limit = 10

	resp, err = client.QueryRecords("your_object_type", query)
	if err != nil {
		log.Fatalf("Error querying records: %v", err)
	}
	fmt.Printf("Query results: %+v\n", resp)

	// Example: Update a record
	updateData := map[string]interface{}{
		"description": "Updated description",
	}

	resp, err = client.UpdateRecord("your_object_type", "record_id", updateData)
	if err != nil {
		log.Fatalf("Error updating record: %v", err)
	}
	fmt.Printf("Updated record: %+v\n", resp)

	// Example: Delete a record
	resp, err = client.DeleteRecord("your_object_type", "record_id")
	if err != nil {
		log.Fatalf("Error deleting record: %v", err)
	}
	fmt.Printf("Deleted record: %+v\n", resp)
}
