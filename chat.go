package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	ChatCompletionURL = "/v1/chat/completions"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
)

func (c *Client) CreateChatCompletion(request *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	res, err := c.DoCommonPostReq(request, c.apiURL+ChatCompletionURL)
	if err != nil {
		return nil, err
	}
	typedRes, ok := res.(*ChatCompletionResponse)
	if !ok {
		return nil, err
	}
	return typedRes, nil
}

func (c *Client) GetChatCompletionStreamReader(request *ChatCompletionRequest) (*bufio.Reader, error) {
	if !request.Stream {
		request.Stream = true
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.apiURL+ChatCompletionURL, bytes.NewBuffer(reqBody))
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

func (c *Client) CreateChatCompletionStream(request *ChatCompletionRequest, onData func(response *ChatCompletionStreamResponse)) error {
	reader, err := c.GetChatCompletionStreamReader(request)
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
		var output *ChatCompletionStreamResponse
		err = json.Unmarshal(line, output)
		if err != nil {
			return err
		}
		onData(output)
	}
	return nil
}
