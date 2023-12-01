package middleware

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func ProductAI(userInput, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "",
		},

		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}
