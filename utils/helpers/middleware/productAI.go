package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type ProductsAI interface {
	GetAllPruducts(ctx echo.Context, id uint) (map[uint]string, error)
	ProductAI(productMap, openAIKey string) (string, error)
}

type ProductAIImpl struct {
	DB *gorm.DB
}

// func (middleware *ProductAIImpl) GetAllPruducts(ctx echo.Context, id uint) (map[uint]string, error) {
// 	products := []domain.Product{}

// 	result := middleware.DB.Preload("ID").Preload("Name").Where("deleted_at IS NULL").Find(&products)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	productMap := make(map[uint]string)
// 	for _, product := range products {
// 		productMap[product.ID] = product.Name
// 	}

// 	return productMap, nil
// }

func ProductAI(productMap map[string]uint, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo

	productMapStr := convertMapToString(productMap)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are the person who works at the cafe. You are a very experienced person in your field. You will be asked to give us one of your best recommendations of ",
		},

		{
			Role:    openai.ChatMessageRoleUser,
			Content: productMapStr,
		},
	}

	resp, err := getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func convertMapToString(productMap map[string]uint) string {
    // Implementasi konversi map menjadi string, contoh:
    var result []string
    for key, value := range productMap {
        result = append(result, fmt.Sprintf("%s:%d", key, value))
    }
    return strings.Join(result, ", ")
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
