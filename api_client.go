package io_base_service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

// APIClient struct to send test API requests
type APIClient struct {
	App *fiber.App
}

// NewAPIClient initializes a test client for Fiber app
func NewAPIClient(app *fiber.App) *APIClient {
	return &APIClient{App: app}
}

// SendRequest sends an HTTP request to the Fiber server
func (client *APIClient) SendRequest(method, url string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}

	req := httptest.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	return client.App.Test(req)
}
