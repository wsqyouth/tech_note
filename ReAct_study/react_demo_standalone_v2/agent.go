package main

import (
	"fmt"
	"strings"
)

// LLM 接口定义了与大语言模型交互的规范
type LLM interface {
	Generate(context string) (string, error)
}

// GoogleLLM 是与 Google LLM API 交互的实现
type GoogleLLM struct {
	// 在实际应用中，这里会包含 API key、HTTP client 等
}

func NewGoogleLLM() *GoogleLLM {
	return &GoogleLLM{}
}

// Generate 方法模拟了调用真实 LLM API 的行为
func (llm *GoogleLLM) Generate(prompt string) (string, error) {
	// 真实场景下，这里会是调用 Google LLM API 的 HTTP 请求。
	// 在本 demo 中，我们通过检查 prompt 的内容来模拟 LLM 在不同阶段的响应。

	// 阶段1：初始状态，历史轨迹为空。LLM 决定需要搜索。
	if strings.Contains(prompt, "无历史执行轨迹") {
		return `思维：好的，任务是“用费曼学习法讲解 Gin 中间件的实现”。首先，我需要一些基础资料来了解 Gin 中间件的核心概念和实现方式。我将使用 google_search 工具来获取这些信息。
行动：google_search[gin framework middleware implementation]`, nil
	}

	// 阶段2：历史轨迹中包含了 google_search 的观察结果。LLM 决定总结并回答。
	if strings.Contains(prompt, "观察：") {
		return "思维：搜索结果提供了足够的信息。我了解了中间件的定义、核心是 c.Next() 和 c.Abort()，以及通过 router.Use() 等方法进行应用的。现在我可以整合这些信息，并使用费曼学习法来构建一个简单易懂的解释，最后完成任务。\n" +
			"行动：finish[当然，让我们用一个简单的比喻来理解 Gin 的中间件。\n\n" +
			"想象一下，你开了一家汉堡店，一个顾客的订单（HTTP 请求）从前台传到后厨。这个订单需要经过好几个岗位才能完成：\n\n" +
			"1.  **配菜员** (第一个中间件)\n" +
			"2.  **烤肉员** (第二个中间件)\n" +
			"3.  **打包员** (主厨，也就是你的最终路由处理器)\n\n" +
			"在 Gin 中，你使用 `router.Use()` 就相当于在你的后厨流水线上增加一个“岗位”（中间件）。\n\n" +
			"每个“岗位”上的员工拿到订单后，会做两件事：\n" +
			"- **处理自己的任务**：比如配菜员准备好生菜和番茄。这对应中间件里 `c.Next()` 之前的代码。\n" +
			"- **喊“下一个！” (`c.Next()`)**: 把订单传给流水线上的下一个人。\n\n" +
			"如果某个员工（比如配菜员）发现订单有问题（比如顾客点的食材没了），他可以大喊一声“别做了！” (`c.Abort()`)，这样后面的烤肉员和打包员就不会再接到这个订单了。\n\n" +
			"所以，Gin 的中间件就是这样一条流水线，请求会按顺序流经你在代码中注册的每一个中间件函数，直到最后一个主处理器。这个“流水线”在 Gin 源码中的核心体现就是 `gin.Engine` 中的 `HandlersChain`，它是一个函数切片（`[]HandlerFunc`），忠实地记录了你要依次执行的所有“岗位”。\n\n" +
			"核心代码位置：你可以在 `github.com/gin-gonic/gin/gin.go` 文件中找到 `Use()` 方法和 `Engine` 结构体中的 `HandlersChain`。]", nil
	}

	// 默认情况，或模拟出现意外
	return `思维：出现了一些意料之外的情况，我需要重新评估当前的状态。
行动：finish[任务失败，无法根据当前信息得出结论。]`, nil
}

// buildPrompt 构建发送给真实 LLM 的提示词
// 这是 ReAct 模式的精髓所在
func buildPrompt(task string, tools []BaseTool, context string) string {
	var toolDescriptions []string
	for _, tool := range tools {
		toolDescriptions = append(toolDescriptions, fmt.Sprintf("- %s: %s", tool.GetName(), tool.GetDescription()))
	}

	prompt := fmt.Sprintf(`
你是一个强大的 AI 助手，遵循 ReAct 模式来解决问题。你的任务是：%s

为了完成任务，你可以使用以下工具：
%s

你的思考和行动过程遵循“思维-行动-观察”的循环。

你需要遵循以下格式：
思维：[你对当前情况的分析和下一步计划]
行动：[你要执行的动作，必须是 'finish[你的最终答案]' 或 '工具名[参数]']

这是你之前的执行历史：
%s

现在，请生成你下一步的思维和行动。
`, task, strings.Join(toolDescriptions, "\n"), context)

	return prompt
}
