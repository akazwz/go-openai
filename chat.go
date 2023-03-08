package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ChatCompletionUrl = "/v1/chat/completions"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
)

func (c *Client) CreateChatCompletion(request *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, c.apiUrl+ChatCompletionUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	var data ChatCompletionResponse
	err = c.doRequest(req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetChatCompletionStreamReader(request *ChatCompletionRequest) (*bufio.Reader, error) {
	if !request.Stream {
		request.Stream = true
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling request: ", err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.apiKey+ChatCompletionUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}
	c.warpRequest(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error doing request: ", err)
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
		line, err := reader.ReadBytes('\n')
		if err != nil {
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
		output := new(ChatCompletionStreamResponse)
		err = json.Unmarshal(line, output)
		if err != nil {
			fmt.Println("Error unmarshalling line: ", err)
			return err
		}
		onData(output)
	}
	return nil
}
