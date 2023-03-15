package prompt

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/sashabaranov/go-openai"
)

var (
	Stream          = os.Getenv("GPT-STREAM") == "true"
	AuthToken       = os.Getenv("GPT-API-KEY")
	MaxTokens       = 1000
	Temperature     = float32(0)
	ChatModel       = openai.GPT3Dot5Turbo
	CompletionModel = openai.GPT3TextDavinci003
)

func AskCompletion(text string) error {
	fmt.Println(text)
	c := openai.NewClient(AuthToken)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:       ChatModel,
		MaxTokens:   MaxTokens,
		Temperature: Temperature,
		Stream:      Stream,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
	}
	if Stream {
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			return fmt.Errorf("Chat Completion Stream error: %v\n", err)
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println()
				return nil
			}
			if err != nil {
				return fmt.Errorf("\nStream error: %v\n", err)
			}
			fmt.Print(resp.Choices[0].Delta.Content)
		}
	} else {
		resp, err := c.CreateChatCompletion(ctx, req)
		if err != nil {
			return fmt.Errorf("Chat Completion error: %v\n", err)
		}
		fmt.Println(resp.Choices[0].Message.Content)
		return nil
	}
}

func CodeCompletion(text string) error {
	fmt.Print(text)
	c := openai.NewClient(AuthToken)
	ctx := context.Background()
	req := openai.CompletionRequest{
		Model:       CompletionModel,
		Prompt:      text,
		MaxTokens:   MaxTokens,
		Temperature: Temperature,
		Stream:      Stream,
	}
	if Stream {
		stream, err := c.CreateCompletionStream(ctx, req)
		if err != nil {
			return fmt.Errorf("Completion Stream error: %v\n", err)
		}
		defer stream.Close()
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println()
				return nil
			}
			if err != nil {
				return fmt.Errorf("\nStream error: %v\n", err)
			}
			fmt.Print(response.Choices[0].Text)
		}
	} else {
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			return fmt.Errorf("\nCompletion error: %v\n", err)
		}
		fmt.Println(resp.Choices[0].Text)
		return nil
	}
}
