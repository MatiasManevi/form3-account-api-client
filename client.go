package form3Client

import (
    "net/http"
    "time"
    "fmt"
	"encoding/json"
	"errors"
	"context"
)

const (
	// Points to form3 account API
	Host = "http://localhost:8080/v1"

	// Used to limit http.Client waiting time.
	httpClientTimeout = 30 * time.Second
)

type Client struct {
	Host       string
	HTTPClient *http.Client
}

type successResponse struct {
	Data interface{} `json:"data"`
	Links interface{} `json:"links"`
}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func NewClient() *Client {
	client := &http.Client{
		Timeout: httpClientTimeout,
	}
    return &Client{
        Host: Host,
        HTTPClient: client,
    }
}

func (c *Client) doRequest(req *http.Request, v interface{}) error {
	// Use context on request to control reuqest deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	// Setting common headers
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    req.Header.Set("Accept", "application/json; charset=utf-8")

	// Run request
	res, err := c.HTTPClient.Do(req)
	// Handle HTTP request errors
    if err != nil {
        return err
    }
    defer res.Body.Close()

	// Checking for errors in response status code
    if res.StatusCode != http.StatusOK {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			// Error response couldn't be decoded
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}
		
		return errors.New(errRes.ErrorMessage)
    }

	if v != nil {
		response := successResponse{
			Data: v,
		}
		
		// Checking for errors in response decoding
		if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
			return err
		}
	}
	
	return nil
}