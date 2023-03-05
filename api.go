package openai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) SetApiKey(apiKey string) {
	c.apiKey = apiKey
}

func (c *Client) SetProxy(proxyURL string) error {
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	c.httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	return nil
}

func (c *Client) UseEnvProxy() {
	c.httpClient.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
}

func (c *Client) warpRequest(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
}

func (c *Client) doRequest(req *http.Request, data any) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	if data != nil {
		err = json.NewDecoder(response.Body).Decode(data)
		if err != nil {
			return err
		}
	}
	return nil
}
