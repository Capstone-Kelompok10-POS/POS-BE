package helpers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/sashabaranov/go-openai"
)

type ProductsAI interface {
	ProductAI(productMap, openAIKey, userInput string) (string, error)
}

type ProductAIImpl struct {
	DB *gorm.DB
}

type ProductDataAIRecommended struct {
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

func ProductAI(productMap map[uint]ProductDataAIRecommended, openAIKey, userInput string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo

	productMapStr := convertMapToString(productMap)
	if productMapStr == "" {
		return "", errors.New("product is empty")
	}
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Anda adalah asisten virtual dalam sistem rekomendasi kafe. Anda adalah orang yang sangat berpengalaman di bidang Anda. Anda akan diminta untuk memberikan rekomendasi terbaik Anda dari semua menu di cafe. Berikan lima rekomendasi terbaik anda jika input meminta makanan maka berikan rekomendasi makanan jika input meminta minuman maka berikan rekomendasi minuman berikut ini adalah product dari cafenya" + productMapStr + "Jika sebelum prompt ini tidak terdapat product yang diberikan maka berikan response product tidak ditemukan",
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

func convertMapToString(productMap map[uint]ProductDataAIRecommended) string {
	// Implementasi konversi map menjadi string, contoh:
	var result []string
	for key, value := range productMap {
		result = append(result, fmt.Sprintf("%d:%s:%s:", key, value.Name, value.Ingredients))
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
