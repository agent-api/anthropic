package anthropic

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/agent-api/anthropic/client"
	"github.com/agent-api/anthropic/models"
	"github.com/agent-api/core"
)

// Provider implements the LLMProvider interface for Anthropic
type Provider struct {
	host string
	port int

	model *core.Model

	// client is the internal Ollama HTTP client
	client *client.AnthropicClient

	logger slog.Logger
}

type ProviderOpts struct {
	BaseURL string
	Port    int
	APIKey  string

	Logger *slog.Logger
}

// NewProvider creates a new Ollama provider
func NewProvider(opts *ProviderOpts) (*Provider, error) {
	ctx := context.Background()
	opts.Logger.Info("Creating new anthropic provider")

	clientOpts := &client.AnthropicClientOpts{
		Model:  models.CLAUDE_3_5_SONNET,
		Logger: opts.Logger,
	}

	if opts.APIKey != "" {
		clientOpts.APIKey = opts.APIKey
	}

	client, err := client.NewClient(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("error creating anthropic client: %w", err)
	}

	return &Provider{
		client: client,
		logger: *opts.Logger,
	}, nil
}

func (p *Provider) GetCapabilities(ctx context.Context) (*core.Capabilities, error) {
	p.logger.Info("Fetching capabilities")

	// Placeholder for future implementation
	p.logger.Info("GetCapabilities method is not implemented yet")

	return nil, nil
}

func (p *Provider) UseModel(ctx context.Context, model *core.Model) error {
	p.logger.Info("Setting model", "modelID", model.ID)

	p.model = model

	return nil
}

// Generate implements the LLMProvider interface for basic responses
func (p *Provider) Generate(ctx context.Context, opts *core.GenerateOptions) (*core.Message, error) {
	p.logger.Info("Generate request received", "modelID", p.model.ID)

	resp, err := p.client.Chat(ctx, &client.ChatRequest{
		Model:    p.model.ID,
		Messages: opts.Messages,
		Tools:    opts.Tools,
	})

	if err != nil {
		p.logger.Error(err.Error(), "Error calling client chat method", err)
		return nil, fmt.Errorf("error calling client chat method: %w", err)
	}

	return &core.Message{
		Role:      core.AssistantMessageRole,
		Content:   resp.Message.Content,
		ToolCalls: resp.Message.ToolCalls,
	}, nil
}

// GenerateStream streams the response token by token
func (p *Provider) GenerateStream(ctx context.Context, opts *core.GenerateOptions) (<-chan *core.Message, <-chan string, <-chan error) {
	p.logger.Info("Starting stream generation", "modelID", p.model.ID)

	msgChan := make(chan *core.Message)
	deltaChan := make(chan string)
	errChan := make(chan error, 1)

	p.logger.Error("stream generation not implemented yet")

	defer close(msgChan)
	defer close(deltaChan)
	defer close(errChan)

	return msgChan, deltaChan, errChan
}
