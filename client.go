package form3Client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Host       string
	HTTPClient *http.Client
}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

type ClientOptions struct {
	Host    string
	Timeout time.Duration
}

func NewClient(options *ClientOptions) *Client {
	// Points to form3 account API
	host := os.Getenv("API_HOST")
	// Used to limit http.Client waiting time.
	timeout := 30 * time.Second

	if options != nil {
		host = options.Host
		timeout = options.Timeout
	}

	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		Host:       host,
		HTTPClient: client,
	}
}

func (c *Client) doRequest(req *http.Request, Response interface{}) error {
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
	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		// Status codes 200 and 201 returned for successful GET-POST requests

		// Checking for errors in response decoding data into go struct
		if err = json.NewDecoder(res.Body).Decode(Response); err != nil {
			return err
		}
	case http.StatusNoContent:
		// Status code 204 returned for successful DELETE requests
		return nil
	case http.StatusInternalServerError:
		// Status code 500 is a server error
		return errors.New("The API service is currently unavailable")
	default:
		// Anything else than a 200/201/204/500
		errRes := errorResponse{}
		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			// Error response couldn't be decoded
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}

		return errors.New(errRes.ErrorMessage)
	}

	return nil
}
