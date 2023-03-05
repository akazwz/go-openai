package examples

import (
	"fmt"

	"github.com/akazwz/go-openai"
)

var client *openai.Client

func init() {
	client = openai.NewClient("sk-xxx")
}

func ChatDemo() {
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

func StreamChatDemo() {
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

func ChatWithProxy() {
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
