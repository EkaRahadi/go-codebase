package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type HttpClient interface {
	Get(ctx context.Context, url string) (string, error)
	GetWithQuery(ctx context.Context, baseURL string, queryParams map[string]string) (string, error)
	GetWithPath(ctx context.Context, baseURL string, path string) (string, error)
	GetWithPathAndQuery(ctx context.Context, baseURL string, path string, queryParams map[string]string) (string, error)
	Post(ctx context.Context, url string, body []byte) (string, error)
	Put(ctx context.Context, url string, body []byte) (string, error)
	Patch(ctx context.Context, url string, body []byte) (string, error)
	Delete(ctx context.Context, url string) (string, error)
}

type Client struct {
	client *http.Client
}

var timeout = 5 * time.Second
var httpClient *Client

func NewClient() HttpClient {
	if httpClient == nil {
		httpClient = &Client{
			client: &http.Client{
				Timeout:   timeout,
				Transport: otelhttp.NewTransport(http.DefaultTransport),
			},
		}
		return httpClient
	}

	return httpClient
}

func (c *Client) Get(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request: %w", err)
	}
	return c.doRequest(req)
}

func (c *Client) GetWithQuery(ctx context.Context, baseURL string, queryParams map[string]string) (string, error) {
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	query := reqURL.Query()
	for key, value := range queryParams {
		query.Set(key, value)
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request: %w", err)
	}

	return c.doRequest(req)
}

func (c *Client) GetWithPath(ctx context.Context, baseURL string, path string) (string, error) {
	fullURL := fmt.Sprintf("%s/%s", baseURL, path)
	reqURL, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request: %w", err)
	}

	return c.doRequest(req)
}

func (c *Client) GetWithPathAndQuery(ctx context.Context, baseURL string, path string, queryParams map[string]string) (string, error) {
	fullURL := fmt.Sprintf("%s/%s", baseURL, path)
	reqURL, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	query := reqURL.Query()
	for key, value := range queryParams {
		query.Set(key, value)
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request: %w", err)
	}

	return c.doRequest(req)
}

func (c *Client) Post(ctx context.Context, url string, body []byte) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

func (c *Client) Put(ctx context.Context, url string, body []byte) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create PUT request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

func (c *Client) Patch(ctx context.Context, url string, body []byte) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create PATCH request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

func (c *Client) Delete(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create DELETE request: %w", err)
	}
	return c.doRequest(req)
}

func (c *Client) doRequest(req *http.Request) (string, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return string(body), fmt.Errorf("received error response: %s", resp.Status)
	}

	return string(body), nil
}
