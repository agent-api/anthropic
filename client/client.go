package client

import (
	"context"
	"log/slog"

	"github.com/agent-api/core/types"

	"github.com/anthropics/anthropic-sdk-go"
)

type AnthropicClient struct {
	//opts []option.RequestOption

	client *anthropic.Client

	model string

	logger *slog.Logger
}

type AnthropicClientOpts struct {
	Logger *slog.Logger
	Model  *types.Model
}

func NewClient(ctx context.Context, opts *AnthropicClientOpts) (*AnthropicClient, error) {
	client := anthropic.NewClient()

	return &AnthropicClient{
		client: client,
		model:  opts.Model.ID,
		logger: opts.Logger,
	}, nil
}

func (c *AnthropicClient) Chat(ctx context.Context, req *ChatRequest) (ChatResponse, error) {
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
		panic(err.Error())
	}

	return ChatResponse{
		Message: types.Message{
			Content: res.Content[0].Text,
		},
	}, nil
}
