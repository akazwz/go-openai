package main

import (
	"fmt"

	"github.com/akazwz/go-openai"
)

func main() {
	//completionDemo()
	// completionStreamDemo()
	GetCredit()
}

func completionDemo() {
	client := openai.NewClient("sk-xxx")
	request := &openai.CompletionRequest{
		Model:  openai.ModelTextDavinci003,
		Prompt: []string{"Hello world"},
	}
	completionResponse, err := client.CreateCompletion(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(completionResponse.Choices[0].Text)
}

func completionStreamDemo() {
	client := openai.NewClient("sk-xxx")
	request := &openai.CompletionRequest{
		Model:  openai.ModelTextDavinci003,
		Prompt: []string{"Hello world"},
	}
	err := client.CreateCompletionStream(request, func(response *openai.CompletionStreamResponse) {
		fmt.Println("response: ", response.Choices[0].Text)
	})
	if err != nil {
		panic(err)
	}
}

func ChatDemo() {
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

func StreamChatDemo() {
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

func ChatWithProxy() {
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

func GetCredit() {
	client := openai.NewClient("sk-xxx")
	creditResponse, err := client.GetCredit()
	if err != nil {
		panic(err)
	}
	fmt.Println(creditResponse.TotalAvailable)
	fmt.Println(creditResponse.TotalGranted)
	fmt.Println(creditResponse.TotalUsed)
}
