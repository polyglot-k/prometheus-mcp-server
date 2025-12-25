package prometheus

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Config holds the configuration for Prometheus client
type Config struct {
	URL           string
	Username      string
	Password      string
	Token         string
	OrgID         string
	SkipSSLVerify bool
}

// Client handles communication with Prometheus API
type Client struct {
	config Config
	client *http.Client
}

// NewClient creates a new Prometheus client
func NewClient() *Client {
	config := Config{
		URL:           os.Getenv("PROMETHEUS_URL"),
		Username:      os.Getenv("PROMETHEUS_USERNAME"),
		Password:      os.Getenv("PROMETHEUS_PASSWORD"),
		Token:         os.Getenv("PROMETHEUS_TOKEN"),
		OrgID:         os.Getenv("ORG_ID"),
		SkipSSLVerify: strings.ToLower(os.Getenv("PROMETHEUS_URL_SSL_VERIFY")) == "false",
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SkipSSLVerify},
	}

	return &Client{
		config: config,
		client: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		},
	}
}

// Request makes a request to Prometheus API
func (c *Client) Request(endpoint string, params map[string]string) (interface{}, error) {
	if c.config.URL == "" {
		return nil, fmt.Errorf("PROMETHEUS_URL not configured")
	}

	baseURL := strings.TrimRight(c.config.URL, "/")
	reqURL, err := url.Parse(fmt.Sprintf("%s/api/v1/%s", baseURL, endpoint))
	if err != nil {
		return nil, err
	}

	q := reqURL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Auth headers
	if c.config.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.Token)
	} else if c.config.Username != "" && c.config.Password != "" {
		req.SetBasicAuth(c.config.Username, c.config.Password)
	}

	if c.config.OrgID != "" {
		req.Header.Set("X-Scope-OrgID", c.config.OrgID)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("prometheus API error: status code %d", resp.StatusCode)
	}

	var result struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
		Error  string      `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("prometheus API error: %s", result.Error)
	}

	return result.Data, nil
}
