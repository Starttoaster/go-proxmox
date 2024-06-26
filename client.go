package proxmox

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://localhost:8006/"
	apiPath        = "api2/json/"
)

// Client for the Proxmox API
type Client struct {
	// HTTP retryable client for the API
	client *http.Client

	// Base URL for API requests. Defaults to https://localhost:8006/,
	// but can be changed to any remote endpoint.
	baseURL *url.URL

	// tokenID is the identifier given for a Proxmox API token
	tokenID string

	// token is the token secret
	token string

	// Services for each resource in the Proxmox API
	Nodes   *NodeService
	Cluster *ClusterService
}

// NewClient returns a new Proxmox API client
func NewClient(tokenID string, token string, options ...ClientOptionFunc) (*Client, error) {
	if token == "" || tokenID == "" {
		return nil, fmt.Errorf("can not create Proxmox API client without a token ID and token")
	}

	c := &Client{
		tokenID: tokenID,
		token:   token,
	}

	// Set the client default fields
	_ = c.setBaseURL(defaultBaseURL)
	_ = c.setHTTPClient(&http.Client{})

	// Apply any given options
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	// Create all the Proxmox API services
	c.Nodes = &NodeService{client: c}
	c.Cluster = &ClusterService{client: c}

	return c, nil
}

// ClientOptionFunc can be used to customize a new Proxmox API client
type ClientOptionFunc func(*Client) error

// WithBaseURL sets the URL for API requests to something other than localhost.
// API path is applied automatically if unspecified.
// Default: "https://localhost:8006/"
func WithBaseURL(urlStr string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// setBaseURL sets the URL for API requests
func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiPath) {
		baseURL.Path += apiPath
	}

	// Update the base URL of the client
	c.baseURL = baseURL

	return nil
}

// WithHTTPClient sets the HTTP client for API requests to something other than the default Go http Client
func WithHTTPClient(client *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		return c.setHTTPClient(client)
	}
}

// setHTTPClient sets the HTTP client for API requests
func (c *Client) setHTTPClient(client *http.Client) error {
	c.client = client
	return nil
}
