package main

import (
	"context"
	"fmt"
	"log"
	//"strings"

	"google.golang.org/genai"
)

// LLM 接口定义了与大语言模型交互的规范
type LLM interface {
	Generate(prompt string) (string, error)
}

// GoogleLLM 是与 Google LLM API 交互的实现
type GoogleLLM struct {
	client *genai.Client
}

// NewGoogleLLM 初始化 GoogleLLM 客户端
// 它会从环境变量 `GEMINI_API_KEY` 自动加载 API 密钥。
func NewGoogleLLM() *GoogleLLM {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		// 如果无法创建客户端（例如，没有设置API密钥），则程序会终止
		log.Fatalf("无法创建 GenAI 客户端，请确保已设置 GEMINI_API_KEY 环境变量: %v", err)
	}
	return &GoogleLLM{client: client}
}

// Generate 调用真实的 Google Gemini API 来生成思维和行动
func (llm *GoogleLLM) Generate(prompt string) (string, error) {
	fmt.Println("\n>>>>>> 正在调用 Google API... <<<<<<")
	ctx := context.Background()

	// 调用 `GenerateContent` 方法
	result, err := llm.client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview", // 使用用户示例中提到的模型 gemini-2.5-flash
		genai.Text(prompt),       // 直接传入 genai.Text(prompt)
		nil,                      // 传入 nil 作为 GenerateContentConfig
	)
	if err != nil {
		return "", fmt.Errorf("调用 GenerateContent API 失败: %w", err)
	}

	// 简化解析响应，直接使用 result.Text()
	return result.Text(), nil
}
