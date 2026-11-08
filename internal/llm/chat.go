// Copyright (c) 2026 heathcetide. All rights reserved.

package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Protocol identifies the wire format.
type Protocol string

const (
	ProtocolOpenAI    Protocol = "openai"    // OpenAI + most OpenAI-compatible gateways
	ProtocolAnthropic Protocol = "anthropic" // Anthropic Messages API
)

// Config is loaded from yaml `llm:` section.
type Config struct {
	Enabled  bool          `yaml:"enabled"`
	Protocol string        `yaml:"protocol"` // openai | anthropic
	APIKey   string        `yaml:"apiKey"`
	BaseURL  string        `yaml:"baseURL"`
	Model    string        `yaml:"model"`
	Timeout  time.Duration `yaml:"timeout"`
}

// Message is a chat turn.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Client talks to an OpenAI-compatible or Anthropic chat endpoint.
type Client struct {
	cfg        Config
	httpClient *http.Client
}

// New builds a client. Returns nil when disabled or incomplete.
func New(cfg Config) *Client {
	if !cfg.Enabled {
		return nil
	}
	if strings.TrimSpace(cfg.APIKey) == "" || strings.TrimSpace(cfg.Model) == "" {
		return nil
	}
	proto := strings.ToLower(strings.TrimSpace(cfg.Protocol))
	if proto == "" {
		proto = string(ProtocolOpenAI)
	}
	cfg.Protocol = proto
	if cfg.Timeout <= 0 {
		cfg.Timeout = 60 * time.Second
	}
	cfg.BaseURL = strings.TrimRight(strings.TrimSpace(cfg.BaseURL), "/")
	if cfg.BaseURL == "" {
		switch Protocol(proto) {
		case ProtocolAnthropic:
			cfg.BaseURL = "https://api.anthropic.com"
		default:
			cfg.BaseURL = "https://api.openai.com/v1"
		}
	}
	return &Client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

// Enabled reports whether the client can make requests.
func (c *Client) Enabled() bool {
	return c != nil && c.cfg.Enabled && c.cfg.APIKey != "" && c.cfg.Model != ""
}

// Chat sends messages and returns the assistant text.
func (c *Client) Chat(ctx context.Context, system string, user string) (string, error) {
	if !c.Enabled() {
		return "", fmt.Errorf("llm client not configured")
	}
	switch Protocol(c.cfg.Protocol) {
	case ProtocolAnthropic:
		return c.chatAnthropic(ctx, system, user)
	default:
		return c.chatOpenAI(ctx, system, user)
	}
}

// ChatVision understands an image via the same chat model (OpenAI/Anthropic multimodal).
// mime eg image/jpeg; b64 is raw base64 without data-url prefix.
func (c *Client) ChatVision(ctx context.Context, system, user, mime, b64 string) (string, error) {
	if !c.Enabled() {
		return "", fmt.Errorf("llm client not configured")
	}
	if strings.TrimSpace(b64) == "" {
		return "", fmt.Errorf("empty image")
	}
	if mime == "" {
		mime = "image/jpeg"
	}
	switch Protocol(c.cfg.Protocol) {
	case ProtocolAnthropic:
		return c.chatAnthropicVision(ctx, system, user, mime, b64)
	default:
		return c.chatOpenAIVision(ctx, system, user, mime, b64)
	}
}

func (c *Client) chatOpenAIVision(ctx context.Context, system, user, mime, b64 string) (string, error) {
	url := c.cfg.BaseURL
	if !strings.HasSuffix(url, "/chat/completions") {
		url = url + "/chat/completions"
	}
	body := map[string]any{
		"model": c.cfg.Model,
		"messages": []any{
			map[string]any{"role": "system", "content": system},
			map[string]any{
				"role": "user",
				"content": []any{
					map[string]any{"type": "text", "text": user},
					map[string]any{
						"type": "image_url",
						"image_url": map[string]any{
							"url": "data:" + mime + ";base64," + b64,
						},
					},
				},
			},
		},
		"temperature": 0.1,
	}
	raw, err := c.doJSON(ctx, http.MethodPost, url, body, map[string]string{
		"Authorization": "Bearer " + c.cfg.APIKey,
	})
	if err != nil {
		return "", err
	}
	var resp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		return "", fmt.Errorf("openai vision decode: %w", err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		return "", fmt.Errorf("openai api: %s", resp.Error.Message)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("openai api: empty choices")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

func (c *Client) chatAnthropicVision(ctx context.Context, system, user, mime, b64 string) (string, error) {
	url := c.cfg.BaseURL
	if !strings.Contains(url, "/v1/messages") {
		url = url + "/v1/messages"
	}
	mediaType := mime
	switch {
	case strings.Contains(mime, "png"):
		mediaType = "image/png"
	case strings.Contains(mime, "gif"):
		mediaType = "image/gif"
	case strings.Contains(mime, "webp"):
		mediaType = "image/webp"
	default:
		mediaType = "image/jpeg"
	}
	body := map[string]any{
		"model":      c.cfg.Model,
		"max_tokens": 2048,
		"system":     system,
		"messages": []any{
			map[string]any{
				"role": "user",
				"content": []any{
					map[string]any{
						"type": "image",
						"source": map[string]any{
							"type":       "base64",
							"media_type": mediaType,
							"data":       b64,
						},
					},
					map[string]any{"type": "text", "text": user},
				},
			},
		},
		"temperature": 0.1,
	}
	raw, err := c.doJSON(ctx, http.MethodPost, url, body, map[string]string{
		"x-api-key":         c.cfg.APIKey,
		"anthropic-version": "2023-06-01",
	})
	if err != nil {
		return "", err
	}
	var resp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		return "", fmt.Errorf("anthropic vision decode: %w", err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		return "", fmt.Errorf("anthropic api: %s", resp.Error.Message)
	}
	var b strings.Builder
	for _, part := range resp.Content {
		if part.Type == "text" || part.Type == "" {
			b.WriteString(part.Text)
		}
	}
	out := strings.TrimSpace(b.String())
	if out == "" {
		return "", fmt.Errorf("anthropic api: empty content")
	}
	return out, nil
}

func (c *Client) chatOpenAI(ctx context.Context, system, user string) (string, error) {
	url := c.cfg.BaseURL
	if !strings.HasSuffix(url, "/chat/completions") {
		url = url + "/chat/completions"
	}
	body := map[string]any{
		"model": c.cfg.Model,
		"messages": []Message{
			{Role: "system", Content: system},
			{Role: "user", Content: user},
		},
		"temperature": 0.2,
	}
	raw, err := c.doJSON(ctx, http.MethodPost, url, body, map[string]string{
		"Authorization": "Bearer " + c.cfg.APIKey,
	})
	if err != nil {
		return "", err
	}
	var resp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		return "", fmt.Errorf("openai decode: %w", err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		return "", fmt.Errorf("openai api: %s", resp.Error.Message)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("openai api: empty choices")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

func (c *Client) chatAnthropic(ctx context.Context, system, user string) (string, error) {
	url := c.cfg.BaseURL
	if !strings.Contains(url, "/v1/messages") {
		url = url + "/v1/messages"
	}
	body := map[string]any{
		"model":      c.cfg.Model,
		"max_tokens": 2048,
		"system":     system,
		"messages": []map[string]string{
			{"role": "user", "content": user},
		},
		"temperature": 0.2,
	}
	raw, err := c.doJSON(ctx, http.MethodPost, url, body, map[string]string{
		"x-api-key":         c.cfg.APIKey,
		"anthropic-version": "2023-06-01",
	})
	if err != nil {
		return "", err
	}
	var resp struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		return "", fmt.Errorf("anthropic decode: %w", err)
	}
	if resp.Error != nil && resp.Error.Message != "" {
		return "", fmt.Errorf("anthropic api: %s", resp.Error.Message)
	}
	var b strings.Builder
	for _, part := range resp.Content {
		if part.Type == "text" || part.Type == "" {
			b.WriteString(part.Text)
		}
	}
	out := strings.TrimSpace(b.String())
	if out == "" {
		return "", fmt.Errorf("anthropic api: empty content")
	}
	return out, nil
}

func (c *Client) doJSON(ctx context.Context, method, url string, body any, headers map[string]string) ([]byte, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(res.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("llm http %d: %s", res.StatusCode, truncate(string(raw), 400))
	}
	return raw, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
