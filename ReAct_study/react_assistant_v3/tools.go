package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// BaseTool 是所有工具都必须实现的通用接口
type BaseTool interface {
	GetName() string
	GetDescription() string
	// Run 执行工具的具体操作，params 是一个字符串，具体的参数格式由每个工具自行定义
	Run(params string) (string, error)
}

// ==================== SaveToFileTool ====================

// SaveToFileTool 用于将内容保存到 Markdown 文件
type SaveToFileTool struct {
	name        string
	description string
}

// NewSaveToFileTool 创建一个保存文件的工具实例
func NewSaveToFileTool() *SaveToFileTool {
	return &SaveToFileTool{
		name:        "save_to_markdown",
		description: "将最终的报告内容保存到一个 Markdown 文件中。参数格式：'文件名|||报告内容'。",
	}
}

func (t *SaveToFileTool) GetName() string {
	return t.name
}

func (t *SaveToFileTool) GetDescription() string {
	return t.description
}

// sanitizeFilename 清理文件名，移除不安全或无效的字符
func sanitizeFilename(filename string) string {
	// 移除路径相关的字符，防止路径遍历
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, "..", "")

	// 定义一个正则表达式，只允许字母、数字、中文、下划线和短横线
	reg := regexp.MustCompile(`[^a-zA-Z0-9\p{Han}_-]+`)
	processed := reg.ReplaceAllString(filename, "")

	// 限制文件名长度
	if len(processed) > 100 {
		processed = processed[:100]
	}

	// 如果处理后文件名为空，则使用默认名
	if processed == "" {
		return "default_report"
	}
	return processed
}

// Run 实现了保存文件的具体逻辑
// 它期望的参数格式是 "文件名|||报告内容"
func (t *SaveToFileTool) Run(params string) (string, error) {
	parts := strings.SplitN(params, "|||", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("SaveToFileTool 参数格式错误，期望是 '文件名|||报告内容'")
	}

	rawFilename := parts[0]
	content := parts[1]

	// 清理文件名并添加 .md 后缀
	filename := sanitizeFilename(rawFilename) + ".md"

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return "", fmt.Errorf("写入文件 '%s' 失败: %w", filename, err)
	}

	successMsg := fmt.Sprintf("报告已成功保存到文件: %s", filename)
	fmt.Println(successMsg)
	return successMsg, nil
}
