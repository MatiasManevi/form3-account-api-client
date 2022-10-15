package form3Client

import (
    "net/http"
    "time"
	"os"
)

type Client struct {
	host       string
	httpClient *http.Client
}
type successResponse struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}
type errorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func NewClient() *Client {
	host := os.Getenv("API_HOST")
	client := &http.Client{
		Timeout: time.Minute
	}
    return &Client{
        BaseURL: host,
        HTTPClient: client
    }
}
