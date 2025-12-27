package main

import (
	"fmt"
	"strings"
)

// ==================== 工具封装层 ====================

// BaseTool 工具接口，定义标准化调用规范
type BaseTool interface {
	GetName() string
	GetDescription() string
	Run(params string) string
}

// ==================== 上下文管理器 ====================

// TAOStep 存储单个 Thought-Action-Observation 三元组
type TAOStep struct {
	Thought     string
	Action      string
	Observation string
}

// ContextManager 上下文管理器，负责存储和裁剪历史轨迹
type ContextManager struct {
	maxLength     int
	taoTrajectory []TAOStep
}

func NewContextManager(maxLength int) *ContextManager {
	return &ContextManager{
		maxLength:     maxLength,
		taoTrajectory: make([]TAOStep, 0),
	}
}

// AddTAO 添加新的 TAO 步骤
func (cm *ContextManager) AddTAO(thought, action, observation string) {
	cm.taoTrajectory = append(cm.taoTrajectory, TAOStep{
		Thought:     thought,
		Action:      action,
		Observation: observation,
	})
}

// GetContextStr 生成模型可理解的上下文字符串
func (cm *ContextManager) GetContextStr() string {
	if len(cm.taoTrajectory) == 0 {
		return "无历史执行轨迹"
	}

	var builder strings.Builder
	for idx, item := range cm.taoTrajectory {
		builder.WriteString(fmt.Sprintf("步骤%d：思维：%s | 行动：%s | 观察：%s\n",
			idx+1, item.Thought, item.Action, item.Observation))
	}
	return builder.String()
}

// ==================== ReAct 核心循环 ====================

// ReactCoreLoop ReAct 核心循环：控制 TAO 迭代流程
func ReactCoreLoop(task string, tools []BaseTool, llm LLM, maxSteps int) (string, string) {
	// 初始化组件
	contextManager := NewContextManager(8000) // 增加上下文长度
	toolMap := make(map[string]BaseTool)
	for _, tool := range tools {
		toolMap[tool.GetName()] = tool
	}

	// 打印任务和工具信息
	var toolDescriptions []string
	for name, tool := range toolMap {
		toolDescriptions = append(toolDescriptions, fmt.Sprintf("- %s: %s", name, tool.GetDescription()))
	}
	fmt.Printf("\n========================================\n")
	fmt.Printf("任务：%s\n", task)
	fmt.Printf("可用工具：\n%s\n", strings.Join(toolDescriptions, "\n"))
	fmt.Printf("========================================\n\n")

	// 循环迭代
	for step := 0; step < maxSteps; step++ {
		// 1. 获取当前上下文
		contextStr := contextManager.GetContextStr()

		// 2. 构建 Prompt 并调用 LLM 生成思维与行动
		prompt := buildPrompt(task, tools, contextStr)
		llmOutput, err := llm.Generate(prompt)
		if err != nil {
			observation := fmt.Sprintf("调用 LLM 失败：%v", err)
			contextManager.AddTAO("错误", "无", observation)
			fmt.Printf("步骤%d：%s\n", step+1, observation)
			continue
		}

		// 3. 解析思维与行动
		var thought, action string
		thoughtAndAction := strings.SplitN(llmOutput, "行动：", 2)
		thought = strings.TrimSpace(strings.TrimPrefix(thoughtAndAction[0], "思维："))
		if len(thoughtAndAction) > 1 {
			action = strings.TrimSpace(thoughtAndAction[1])
		}

		if thought == "" || action == "" {
			observation := "解析失败：无法从 LLM 输出中提取有效的思维或行动"
			contextManager.AddTAO(thought, action, observation)
			fmt.Printf("步骤%d：%s\n", step+1, observation)
			continue
		}

		// 4. 执行行动并获取观察结果
		var observation string
		if strings.HasPrefix(action, "finish[") {
			// 任务完成，提取结果
			result := strings.TrimPrefix(action, "finish[")
			if lastBracket := strings.LastIndex(result, "]"); lastBracket != -1 {
				result = result[:lastBracket]
			}

			// 打印最终的思考和行动
			fmt.Printf("步骤%d：思维：%s\n", step+1, thought)
			fmt.Printf("      行动：%s\n", action)
			fmt.Printf("----------------------------------------\n")
			fmt.Println("✅ 任务完成")
			return result, contextManager.GetContextStr()
		}

		// 解析工具名称和参数
		executed := false
		for toolName, tool := range toolMap {
			if strings.HasPrefix(action, toolName+"[") {
				paramStart := len(toolName) + 1
				paramEnd := strings.LastIndex(action, "]")
				if paramEnd > paramStart {
					params := action[paramStart:paramEnd]
					observation = tool.Run(params)
					executed = true
					break
				}
			}
		}

		if !executed {
			validTools := make([]string, 0, len(toolMap))
			for name := range toolMap {
				validTools = append(validTools, name)
			}
			observation = fmt.Sprintf("无效行动：%s，支持的工具为 %v", action, validTools)
		}

		// 5. 更新上下文并打印
		contextManager.AddTAO(thought, action, observation)
		fmt.Printf("步骤%d：思维：%s\n", step+1, thought)
		fmt.Printf("      行动：%s\n", action)
		fmt.Printf("      观察：%s\n", strings.TrimSpace(observation))
		fmt.Printf("----------------------------------------\n")
	}

	// 超时终止
	fmt.Println("❌ 任务未完成（已达最大步数）")
	return fmt.Sprintf("任务未完成（已达最大步数%d）", maxSteps), contextManager.GetContextStr()
}

// ==================== 主程序 ====================

func main() {
	// 初始化工具
	tools := []BaseTool{
		NewGoogleSearchTool(),
	}

	// 初始化 LLM
	llm := NewGoogleLLM()

	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          Gin 源码学习助手 (ReAct Go 实现)                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// 定义任务
	task := "请用费曼学习法讲解一下 Gin 框架的中间件是如何实现的，并给出核心代码位置。"

	// 启动 ReAct 循环
	result, _ := ReactCoreLoop(task, tools, llm, 6)

	fmt.Printf("\n\n================ 最终回答 ==================\n")
	fmt.Println(result)
	fmt.Printf("============================================\n")
}
