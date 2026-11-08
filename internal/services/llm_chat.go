// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/llm"
)

// ChatAttachment is an optional image/file for the admin LLM harness.
type ChatAttachment struct {
	Mime     string `json:"mime"`
	Base64   string `json:"base64"`
	FileName string `json:"fileName,omitempty"`
}

// AdminChatRequest is a frontend harness turn.
type AdminChatRequest struct {
	System      string           `json:"system"`
	Message     string           `json:"message"`
	History     []llm.Message    `json:"history"`
	Attachments []ChatAttachment `json:"attachments"`
}

// AdminChatResponse returns model reply for the harness UI.
type AdminChatResponse struct {
	Reply   string `json:"reply"`
	Model   string `json:"model,omitempty"`
	UsedVision bool `json:"usedVision"`
}

// AdminLLMChat runs a single harness turn (text or multimodal).
func AdminLLMChat(ctx context.Context, chat *llm.Client, modelName string, req AdminChatRequest) (*AdminChatResponse, error) {
	if chat == nil || !chat.Enabled() {
		return nil, fmt.Errorf("LLM 未启用或未配置")
	}
	msg := strings.TrimSpace(req.Message)
	if msg == "" && len(req.Attachments) == 0 {
		return nil, fmt.Errorf("请输入内容或上传图片")
	}
	system := strings.TrimSpace(req.System)
	if system == "" {
		system = "你是高原沿线传染病监测系统的助手。回答简洁、可执行；涉及诊断时强调仅供辅助、需人工确认。"
	}
	// Fold short history into user prompt (stateless harness).
	var hist strings.Builder
	for _, h := range req.History {
		role := h.Role
		if role == "" {
			continue
		}
		hist.WriteString(role)
		hist.WriteString(": ")
		hist.WriteString(h.Content)
		hist.WriteString("\n")
	}
	user := msg
	if hist.Len() > 0 {
		user = "对话历史：\n" + hist.String() + "\n用户：" + msg
	}

	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	var imgs []ChatAttachment
	var skipped []string
	for _, a := range req.Attachments {
		m := strings.ToLower(strings.TrimSpace(a.Mime))
		if strings.TrimSpace(a.Base64) == "" {
			continue
		}
		if !strings.HasPrefix(m, "image/") {
			continue
		}
		// MiniMax / many vision APIs reject SVG / animated GIF / HEIC.
		if strings.Contains(m, "svg") || m == "image/svg+xml" || m == "image/heic" || m == "image/heif" {
			name := a.FileName
			if name == "" {
				name = m
			}
			skipped = append(skipped, name)
			continue
		}
		imgs = append(imgs, a)
	}
	if len(skipped) > 0 {
		user += "\n（已跳过不支持的图片格式：" + strings.Join(skipped, "、") + "；请改用 JPG/PNG/WebP）"
	}
	if len(imgs) > 0 {
		// Use first image for vision; mention extra count in text.
		if len(imgs) > 1 {
			user += fmt.Sprintf("\n（另有 %d 张附件图未逐张传入，请主要依据第一张）", len(imgs)-1)
		}
		mime := imgs[0].Mime
		if mime == "" {
			mime = "image/jpeg"
		}
		reply, err := chat.ChatVision(ctx, system, user, mime, imgs[0].Base64)
		if err != nil {
			return nil, err
		}
		return &AdminChatResponse{Reply: reply, Model: modelName, UsedVision: true}, nil
	}
	if len(skipped) > 0 && msg == "" {
		return nil, fmt.Errorf("图片格式不支持（如 SVG），请改用 JPG/PNG/WebP")
	}

	// Non-image attachments: only pass filename hints.
	for _, a := range req.Attachments {
		if strings.TrimSpace(a.FileName) != "" {
			user += "\n[附件] " + a.FileName
		}
	}
	reply, err := chat.Chat(ctx, system, user)
	if err != nil {
		return nil, err
	}
	return &AdminChatResponse{Reply: reply, Model: modelName, UsedVision: false}, nil
}
