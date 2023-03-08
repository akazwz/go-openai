package openai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	CreateImageUrl    = "/v1/images/generations"
	EditImageUrl      = "/v1/images/edits"
	VariationImageUrl = "/v1/images/variations"
)

func (c *Client) CreateImage(request *CreateImageRequest) (*ImageResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, c.apiUrl+CreateImageUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	var data ImageResponse
	err = c.doRequest(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) EditImage(request *EditImageRequest) (*ImageResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, c.apiUrl+EditImageUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	var data ImageResponse
	err = c.doRequest(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) VariationImage(request *VariationImageRequest) (*ImageResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, c.apiUrl+VariationImageUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	var data ImageResponse
	err = c.doRequest(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
