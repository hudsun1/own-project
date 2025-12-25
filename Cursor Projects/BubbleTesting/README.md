# Bubble.io API Testing Project

This project provides a Go client for interacting with the Bubble.io API. It includes functionality for creating, reading, updating, and deleting records in your Bubble.io application.

## Prerequisites

- Go 1.21 or later
- A Bubble.io account and application
- Bubble.io API key

## Setup

1. Clone the repository
2. Set up your environment variables:
   ```bash
   export BUBBLE_API_KEY="your_api_key"
   export BUBBLE_APP_NAME="your_app_name"
   export BUBBLE_BASE_URL="https://api.bubble.io/v1"  # Optional, defaults to this value
   ```

## Project Structure

```
BubbleTesting/
├── api/
│   └── client.go      # API client implementation
├── config/
│   └── config.go      # Configuration management
├── models/
│   └── bubble_data.go # Data structures
├── main.go           # Example usage
├── go.mod           # Go module file
└── README.md        # This file
```

## Usage

The project provides a client for interacting with Bubble.io's API. Here's how to use it:

```go
// Load configuration
cfg, err := config.LoadConfig()
if err != nil {
    log.Fatal(err)
}

// Create API client
client := api.NewClient(cfg)

// Create a record
data := map[string]interface{}{
    "name": "Test Record",
    "description": "This is a test record",
}
resp, err := client.CreateRecord("your_object_type", data)

// Query records
query := models.NewBubbleQuery()
query.AddConstraint("name", "equals", "Test Record")
resp, err = client.QueryRecords("your_object_type", query)

// Update a record
updateData := map[string]interface{}{
    "description": "Updated description",
}
resp, err = client.UpdateRecord("your_object_type", "record_id", updateData)

// Delete a record
resp, err = client.DeleteRecord("your_object_type", "record_id")
```

## Features

- Create, read, update, and delete records
- Query records with constraints and sorting
- Error handling and response parsing
- Configuration management
- Type-safe data structures

## Error Handling

The client includes comprehensive error handling for API requests. All methods return both a response and an error, which should be checked:

```go
resp, err := client.CreateRecord("object_type", data)
if err != nil {
    log.Printf("Error: %v", err)
    return
}
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 