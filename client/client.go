package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/agent-api/core"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AnthropicClient struct {
	//opts []option.RequestOption

	client *anthropic.Client

	model string

	logger *slog.Logger
}

type AnthropicClientOpts struct {
	Logger *slog.Logger
	Model  *core.Model
	APIKey string
}

func NewClient(ctx context.Context, opts *AnthropicClientOpts) (*AnthropicClient, error) {
	requestOpts := []option.RequestOption{}

	if opts.APIKey != "" {
		requestOpts = append(requestOpts, option.WithAPIKey(opts.APIKey))
	}

	client := anthropic.NewClient(requestOpts...)

	return &AnthropicClient{
		client: client,
		model:  opts.Model.ID,
		logger: opts.Logger,
	}, nil
}

func (c *AnthropicClient) Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	// TODO - need to handle adding to anthropic history
	//anthropicMessages := []*googGenAI.Content{
	//{
	//Parts: []googGenAI.Part{
	//googGenAI.Text(req.Messages[0].Content),
	//},
	//Role: "user",
	//},
	//}

	// TODO - need to handle multiple messages

	// TODO - need to handle tools

	res, err := c.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.F(c.model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(req.Messages[0].Content)),
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("error calling anthropic client: %w", err)
	}

	return &ChatResponse{
		Message: core.Message{
			Content: res.Content[0].Text,
		},
	}, nil
}
