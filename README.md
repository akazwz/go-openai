[中文](README_zh.md) | [English](README.md)

# Go Openai Api Wrapper

This is a Go wrapper for the OpenAI API. It is a work in progress and is not yet complete.

## installation

```bash
go get github.com/akazwz/go-openai
```

## features

### 1. gpt-3.5-turbo model support (this model is cheaper 90% than gpt-3, recommend to use this model)

### 2. chat completion support

### 3. chat completion stream support

### 4. proxy support

## chat simple usage

```go
package main

import (
	"fmt"
	"github.com/akazwz/go-openai"
)

func main() {
	client := openai.NewClient("sk-xxx")
	request := &openai.ChatCompletionRequest{
		Model: openai.ModelGPT3Dot5Turbo,
		Messages: []openai.Message{
			{Role: openai.ChatMessageRoleUser, Content: "Hello"},
		},
	}
	completionResponse, err := client.CreateChatCompletion(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(completionResponse.Choices[0].Message.Content)
}
```

## chat stream usage

```go
package main

import (
	"fmt"
	"github.com/akazwz/go-openai"
)

func main() {
	client := openai.NewClient("sk-xxx")
	request := &openai.ChatCompletionRequest{
		Model: openai.ModelGPT3Dot5Turbo,
		Messages: []openai.Message{
			{Role: openai.ChatMessageRoleUser, Content: "Hello"},
		},
	}
	err := client.CreateChatCompletionStream(request, func(response *openai.ChatCompletionStreamResponse) {
		fmt.Println("response: ", response.Choices[0].Delta)
	})
	if err != nil {
		panic(err)
	}
}
```

## usage with proxy

```go
package main

import (
	"fmt"
	"github.com/akazwz/go-openai"
)

func main() {
	client := openai.NewClient("sk-xxx")
	err := client.SetProxy("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}

	request := &openai.ChatCompletionRequest{
		Model: openai.ModelGPT3Dot5Turbo,
		Messages: []openai.Message{
			{Role: openai.ChatMessageRoleUser, Content: "Hello"},
		},
	}

	completionResponse, err := client.CreateChatCompletion(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(completionResponse.Choices[0].Message.Content)
}
```