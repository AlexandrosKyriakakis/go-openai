package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	imagePath := "path/to/img.png"

	// Read the image file
	imgData, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		os.Exit(1)
	}

	// Encode to base64
	base64Str := base64.StdEncoding.EncodeToString(imgData)
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4VisionPreview,
		MaxTokens: 4096,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: []openai.ChatMessageContent{
					{
						Type: openai.ChatMessageContentTypeText,
						Text: "What's in this image",
					},
					{
						Type: openai.ChatMessageContentTypeImage,
						ImageURL: &openai.ChatMessageImageURL{
							Detail: openai.ImageURLDetailHigh,
							URL:    "data:image/png;base64," + base64Str,
						},
					},
				},
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
}
