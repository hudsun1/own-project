package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"BubbleTesting/config"
	"BubbleTesting/models"
)

// Client represents a Bubble.io API client
type Client struct {
	config     *config.Config
	httpClient *http.Client
}

// NewClient creates a new Bubble.io API client
func NewClient(cfg *config.Config) *Client {
	return &Client{
		config: cfg,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

// CreateRecord creates a new record in Bubble.io
func (c *Client) CreateRecord(objectType string, data map[string]interface{}) (*models.BubbleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", c.config.BubbleBaseURL, c.config.BubbleAppName, objectType)

	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.BubbleAPIKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	return c.parseResponse(resp.Body)
}

// QueryRecords queries records from Bubble.io
func (c *Client) QueryRecords(objectType string, query *models.BubbleQuery) (*models.BubbleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", c.config.BubbleBaseURL, c.config.BubbleAppName, objectType)

	body, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error marshaling query: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.BubbleAPIKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	return c.parseResponse(resp.Body)
}

// UpdateRecord updates an existing record in Bubble.io
func (c *Client) UpdateRecord(objectType, recordID string, data map[string]interface{}) (*models.BubbleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.config.BubbleBaseURL, c.config.BubbleAppName, objectType, recordID)

	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %v", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.BubbleAPIKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	return c.parseResponse(resp.Body)
}

// DeleteRecord deletes a record from Bubble.io
func (c *Client) DeleteRecord(objectType, recordID string) (*models.BubbleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.config.BubbleBaseURL, c.config.BubbleAppName, objectType, recordID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.BubbleAPIKey))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	return c.parseResponse(resp.Body)
}

// parseResponse parses the response body into a BubbleResponse
func (c *Client) parseResponse(body io.Reader) (*models.BubbleResponse, error) {
	var response models.BubbleResponse
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &response, nil
}
