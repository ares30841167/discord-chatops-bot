package gitlab

import (
	"errors"
	"net/http"
	"net/url"
)

type Client struct {
	client  *http.Client
	baseURL *url.URL
}

// Initial a new http client instance
func NewClient(baseURL *url.URL) (*Client, error) {
	if baseURL == nil {
		return nil, errors.New("No base URL passed in")
	}
	return &Client{
		client:  &http.Client{},
		baseURL: baseURL,
	}, nil
}

// Send a post request with form to the base url in the client
func (c *Client) PostForm(variables map[string]string) (resp *http.Response, err error) {
	formValues := url.Values{}
	// Fill the variables into url.Values object
	for k, v := range variables {
		formValues.Add(k, v)
	}
	return c.client.PostForm(c.baseURL.Path, formValues)
}
