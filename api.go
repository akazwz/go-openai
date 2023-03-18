package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
	apiURL     string
}

func NewClientWithConfig(config *Client) *Client {
	return &Client{
		apiKey:     config.apiKey,
		httpClient: config.httpClient,
		apiURL:     config.apiURL,
	}
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
		apiURL:     "https://api.openai.com",
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

func (c *Client) SetApiURL(apiURL string) {
	c.apiURL = apiURL
}

func (c *Client) UseEnvProxy() {
	c.httpClient.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
}

func (c *Client) warpRequest(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
}

func (c *Client) doRequest(req *http.Request, data any) error {
	c.warpRequest(req)
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

func (c *Client) DoCommonPostReq(request any, url string) (any, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	var data ImageResponse
	err = c.doRequest(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
