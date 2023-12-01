package middleware

import (
	"context"
	"qbills/models/domain"

	"github.com/jinzhu/gorm"
	"github.com/sashabaranov/go-openai"
)

type MiddlewareImpl struct {
	DB *gorm.DB
}

func (middleware *MiddlewareImpl) GetAllPruducts() (map[uint]string, error) {
	products := []domain.Product{}

	result := middleware.DB.Preload("Name").Where("deleted_at IS NULL").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	productMap := make(map[uint]string)
	for _, product := range products {
		productMap[product.ID] = product.Name
	}

	return productMap, nil
}

func ProductAI(mapproduct, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Give one product recommendation from:",
		},

		{
			Role:    openai.ChatMessageRoleUser,
			Content: mapproduct,
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
