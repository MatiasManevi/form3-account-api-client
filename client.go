package form3Client

import (
    "net/http"
    "time"
	"os"
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
	host := os.Getenv("API_HOST")
	client := &http.Client{
		Timeout: time.Minute,
	}
    return &Client{
        Host: host,
        HTTPClient: client,
    }
}
