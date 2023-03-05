[中文](README_zh.md) | [English](README.md)

# Go Openai Api Wrapper

这是一个Go语言的OpenAI API的包装器。它是一个工作中的项目，还没有完成。

## 安装

```bash
go get github.com/akazwz/go-openai
```

## 特性

### 1. gpt-3.5-turbo 模型支持 (这个模型比gpt-3便宜90%，推荐使用这个模型)

### 2. 聊天自动补全支持

### 3. 聊天自动补全流支持

### 4. 代理支持

## 聊天自动补全简单使用

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

## 聊天自动补全流使用

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

## 使用代理

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

