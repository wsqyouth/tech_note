package main

import "fmt"

// GoogleSearchTool 使用 Google 搜索信息的工具
type GoogleSearchTool struct {
	name        string
	description string
}

func NewGoogleSearchTool() *GoogleSearchTool {
	return &GoogleSearchTool{
		name:        "google_search",
		description: "用于在 Google 上搜索信息，参数为搜索关键词",
	}
}

func (t *GoogleSearchTool) GetName() string {
	return t.name
}

func (t *GoogleSearchTool) GetDescription() string {
	return t.description
}

// Run 实际执行搜索的地方。
// 注意：在当前这个模拟环境中，我们无法真正调用外部 API。
// 我们将在这里放置一个占位符，并返回一个模拟的搜索结果。
// 在下一步与真实 LLM 对接时，这里会被替换为真实的 API 调用。
func (t *GoogleSearchTool) Run(params string) string {
	// 在这个演示中，我们用一个真实的搜索结果来填充，以模拟真实调用
	if params == "gin framework middleware implementation" {
		return `
		Gin 中间件核心要点:
		1. 作用: 中间件用于在主路由处理器前后执行代码，适合日志、认证、错误处理等。
		2. 定义: 一个中间件就是一个函数，其签名为 func(c *gin.Context)。
		3. 核心函数 c.Next(): 在中间件中，调用 c.Next() 会将控制权交给处理链中的下一个函数（另一个中间件或主处理器）。c.Next() 之前的代码在请求到达主处理器前执行，之后的代码在主处理器执行完毕后执行。
		4. 中断函数 c.Abort(): 可以使用 c.Abort() 或 c.AbortWithStatus() 来立即停止处理链，这在认证失败等场景下很有用。
		5. 使用方式:
		   - 全局使用: router.Use(MyMiddleware()) 会让所有请求都经过该中间件。
		   - 路由组使用: router.Group("/admin").Use(AuthMiddleware()) 只对特定分组的路由生效。
		   - 单个路由使用: router.GET("/ping", MyMiddleware(), handler) 只对该路由生效。
		`
	}

	return fmt.Sprintf("没有找到关于 '%s' 的相关信息。", params)
}
