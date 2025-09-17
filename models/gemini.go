package models

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type Gemini struct {
	Message string
	System  string
}

func NewGemini(system, message string) *Gemini {
	return &Gemini{
		Message: message,
		System:  system,
	}
}

func (g *Gemini) SendMessage() (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Models.GenerateContent(ctx, "gemini-pro", genai.Text(g.System+"\n"+g.Message), nil)
	if err != nil {
		return "", err
	}

	if resp.Candidates[0].Content != nil {
		return resp.Text(), nil
	}

	return "", fmt.Errorf("no content found in response")
}
