package form3Client

import (
    "net/url"
    "net/http"
    "io"
)

func buildRequest(method string, path string, data io.Reader) (*http.Request, error) {
	// Checks whether the url path is well formed and valid
	uri, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Creates request
	req, err := http.NewRequest(method, uri.String(), data)
	
	if err != nil {
		return nil, err
	}
	
	return req, nil
}