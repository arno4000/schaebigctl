package ai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func GenAIResponse(promt string, question string) (string, error) {
	chatGPTInput := parseMessage(promt, question)
	client := openai.NewClient(viper.GetString("chatGPTToken"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: chatGPTInput,
				},
			},
		},
	)

	if err != nil {
		return "schaebig 🤣", err
	}

	return resp.Choices[0].Message.Content, err
}

func parseMessage(promt string, question string) string {
	return fmt.Sprintf("%s %s", promt, question)
}
