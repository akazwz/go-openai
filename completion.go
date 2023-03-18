package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	CompletionURL = "/v1/completions"
)

func (c *Client) CreateCompletion(request *CompletionRequest) (*CompletionResponse, error) {
	res, err := c.DoCommonPostReq(request, c.apiURL+CompletionURL)
	if err != nil {
		return nil, err
	}
	typedRes, ok := res.(*CompletionResponse)
	if !ok {
		return nil, err
	}
	return typedRes, nil
}

func (c *Client) GetCompletionStreamReader(request *CompletionRequest) (*bufio.Reader, error) {
	if !request.Stream {
		request.Stream = true
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.apiURL+CompletionURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	c.warpRequest(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(response.Body)
	return reader, nil
}

func (c *Client) CreateCompletionStream(request *CompletionRequest, onData func(response *CompletionStreamResponse)) error {
	reader, err := c.GetCompletionStreamReader(request)
	if err != nil {
		return err
	}
	for {
		line, readErr := reader.ReadBytes('\n')
		if readErr != nil {
			return err
		}
		line = bytes.TrimSpace(line)
		if !bytes.HasPrefix(line, DataPrefix) {
			continue
		}
		line = bytes.TrimPrefix(line, DataPrefix)
		if bytes.HasPrefix(line, DoneSequence) {
			break
		}
		var output *CompletionStreamResponse
		err = json.Unmarshal(line, output)
		if err != nil {
			return err
		}
		onData(output)
	}
	return nil
}
