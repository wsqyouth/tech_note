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

// FlightSearchTool 航班查询工具
type FlightSearchTool struct {
	name        string
	description string
}

func NewFlightSearchTool() *FlightSearchTool {
	return &FlightSearchTool{
		name:        "flight_search",
		description: "用于查询指定条件的航班信息，参数格式为'出发地,目的地,日期,时段'，时段支持'上午/下午/晚上'",
	}
}

func (t *FlightSearchTool) GetName() string {
	return t.name
}

func (t *FlightSearchTool) GetDescription() string {
	return t.description
}

func (t *FlightSearchTool) Run(params string) string {
	// 模拟航班查询逻辑
	flightMap := map[string]string{
		"深圳,海南,明天,晚上": "符合条件航班列表：1. HU7089（深圳宝安→海口美兰，20:15-21:45，票价480元）；2. CZ6753（深圳宝安→三亚凤凰，21:30-23:05，票价620元）；3. MU2478（深圳宝安→海口美兰，19:40-21:10，票价550元）",
		"深圳,广州,昨天,上午": "符合条件航班列表：1. CZ3201（深圳宝安→广州白云，08:30-09:10，票价230元）；2. HU7125（深圳宝安→广州白云，09:40-10:20，票价280元）",
		"北京,上海,今天,下午": "符合条件航班列表：1. CA1234（北京首都→上海虹桥，14:00-16:30，票价850元）；2. MU5678（北京大兴→上海浦东，15:20-17:50，票价780元）",
	}

	if result, ok := flightMap[params]; ok {
		return result
	}
	return fmt.Sprintf("未检索到相关航班信息（参数：%s）", params)
}

// FlightBookTool 航班预订工具
type FlightBookTool struct {
	name        string
	description string
}

func NewFlightBookTool() *FlightBookTool {
	return &FlightBookTool{
		name:        "flight_book",
		description: "用于预订指定航班，参数格式为'航班号,乘客姓名,身份证号'",
	}
}

func (t *FlightBookTool) GetName() string {
	return t.name
}

func (t *FlightBookTool) GetDescription() string {
	return t.description
}

func (t *FlightBookTool) Run(params string) string {
	// 解析参数
	parts := strings.Split(params, ",")
	if len(parts) < 3 {
		return fmt.Sprintf("航班预订失败：参数格式错误，需要'航班号,乘客姓名,身份证号'")
	}

	flightNo := strings.TrimSpace(parts[0])
	name := strings.TrimSpace(parts[1])
	idCard := strings.TrimSpace(parts[2])

	// 模拟预订成功
	if len(idCard) >= 4 {
		lastFour := idCard[len(idCard)-4:]
		return fmt.Sprintf("航班预订成功：航班号%s，乘客%s（身份证号：%s），请携带有效证件提前2小时到机场办理登机手续", flightNo, name, lastFour)
	}
	return fmt.Sprintf("航班预订成功：航班号%s，乘客%s", flightNo, name)
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

// AddTAO 添加新的 TAO 步骤并裁剪上下文
func (cm *ContextManager) AddTAO(thought, action, observation string) {
	cm.taoTrajectory = append(cm.taoTrajectory, TAOStep{
		Thought:     thought,
		Action:      action,
		Observation: observation,
	})
	cm.pruneTrajectory()
}

// pruneTrajectory 裁剪超长轨迹：保留近期3轮 + 早期摘要
func (cm *ContextManager) pruneTrajectory() {
	trajectoryStr := cm.GetContextStr()
	if len(trajectoryStr) <= cm.maxLength {
		return
	}

	// 保留近期3轮完整轨迹
	trajLen := len(cm.taoTrajectory)
	var recentTrajectory []TAOStep
	if trajLen >= 3 {
		recentTrajectory = cm.taoTrajectory[trajLen-3:]
	} else {
		recentTrajectory = cm.taoTrajectory
	}

	// 生成早期轨迹摘要
	var earlyActions []string
	for i := 0; i < trajLen-3 && i < 2; i++ {
		earlyActions = append(earlyActions, cm.taoTrajectory[i].Action)
	}

	var successObs []string
	for i := 0; i < trajLen-3; i++ {
		if strings.Contains(cm.taoTrajectory[i].Observation, "成功") {
			obs := cm.taoTrajectory[i].Observation
			if len(obs) > 30 {
				obs = obs[:30]
			}
			successObs = append(successObs, obs)
			if len(successObs) >= 1 {
				break
			}
		}
	}

	earlySummary := fmt.Sprintf("早期行动：%s... 关键结果：%v", strings.Join(earlyActions, ", "), successObs)

	// 重构上下文
	cm.taoTrajectory = append([]TAOStep{{
		Thought:     "【早期轨迹摘要】",
		Action:      "",
		Observation: earlySummary,
	}}, recentTrajectory...)
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

// ==================== LLM 模拟器 ====================

// SimulatedLLM 模拟 LLM 输出，根据任务和步骤返回预设的思维和行动
type SimulatedLLM struct {
	task string
	step int
}

func NewSimulatedLLM(task string) *SimulatedLLM {
	return &SimulatedLLM{
		task: task,
		step: 0,
	}
}

// Generate 生成思维和行动（模拟 LLM 的推理过程）
func (llm *SimulatedLLM) Generate(context string) string {
	defer func() { llm.step++ }()

	// 根据任务类型和步骤返回不同的输出
	if strings.Contains(llm.task, "深圳到海南") && strings.Contains(llm.task, "晚上") {
		return llm.generateGoodExample()
	} else if strings.Contains(llm.task, "北京到火星") {
		return llm.generateBadExample()
	}

	// 默认响应
	return "思维：任务无法理解，需要更多信息\n行动：finish[无法完成任务]"
}

// generateGoodExample 好的示例：成功完成任务
func (llm *SimulatedLLM) generateGoodExample() string {
	switch llm.step {
	case 0:
		return `思维：当前任务是查询明天从深圳到海南的航班，选最便宜、航班时间在晚上的那班并预订，无历史数据。需先调用航班查询工具，参数为"深圳,海南,明天,晚上"，获取符合条件的航班列表
行动：flight_search[深圳,海南,明天,晚上]`
	case 1:
		return `思维：已获取明天深圳到海南晚上的航班列表，从观察结果可知最便宜的是HU7089（票价480元）。下一步需调用航班预订工具，参数包含航班号HU7089、乘客信息（假设乘客为李四，身份证号440301199505056789）
行动：flight_book[HU7089,李四,440301199505056789]`
	case 2:
		return `思维：已完成最便宜晚上航班的查询与预订，观察结果显示预订成功，所有任务目标均达成，可提交最终结果
行动：finish[明天从深圳到海南最便宜的晚上航班为HU7089（深圳宝安→海口美兰，20:15-21:45，票价480元），已完成预订，乘客李四（身份证号：6789）]`
	default:
		return `思维：任务已完成，无需进一步行动
行动：finish[任务已完成]`
	}
}

// generateBadExample 坏的示例：无效工具调用、参数错误
func (llm *SimulatedLLM) generateBadExample() string {
	switch llm.step {
	case 0:
		return `思维：需要查询北京到火星的航班信息，调用航班查询工具
行动：flight_search[北京,火星,明天,上午]`
	case 1:
		return `思维：查询失败，尝试使用不存在的工具
行动：rocket_search[北京,火星]`
	case 2:
		return `思维：工具调用失败，尝试直接预订
行动：flight_book[MARS001]`
	default:
		return `思维：多次尝试失败，任务无法完成
行动：finish[无法完成任务：不存在北京到火星的航班]`
	}
}

// ==================== ReAct 核心循环 ====================

// ReactCoreLoop ReAct 核心循环：控制 TAO 迭代流程
func ReactCoreLoop(task string, tools []BaseTool, maxSteps int) (string, string) {
	// 初始化组件
	contextManager := NewContextManager(4000)
	toolMap := make(map[string]BaseTool)
	for _, tool := range tools {
		toolMap[tool.GetName()] = tool
	}

	// 初始化模拟 LLM
	llm := NewSimulatedLLM(task)

	// 构建工具描述
	var toolDescriptions []string
	for name, tool := range toolMap {
		toolDescriptions = append(toolDescriptions, fmt.Sprintf("- %s：%s", name, tool.GetDescription()))
	}

	fmt.Printf("\n========================================\n")
	fmt.Printf("任务：%s\n", task)
	fmt.Printf("可用工具：\n%s\n", strings.Join(toolDescriptions, "\n"))
	fmt.Printf("========================================\n\n")

	// 循环迭代
	for step := 0; step < maxSteps; step++ {
		// 1. 获取当前上下文
		context := contextManager.GetContextStr()

		// 2. 调用模拟 LLM 生成思维与行动
		llmOutput := llm.Generate(context)

		// 3. 解析思维与行动
		lines := strings.Split(llmOutput, "\n")
		var thought, action string
		for _, line := range lines {
			if strings.HasPrefix(line, "思维：") {
				thought = strings.TrimPrefix(line, "思维：")
			} else if strings.HasPrefix(line, "行动：") {
				action = strings.TrimPrefix(line, "行动：")
			}
		}

		if thought == "" || action == "" {
			observation := "解析失败：无法提取思维或行动"
			contextManager.AddTAO(thought, action, observation)
			fmt.Printf("步骤%d：%s\n", step+1, observation)
			continue
		}

		// 4. 执行行动并获取观察结果
		var observation string
		if strings.HasPrefix(action, "finish[") {
			// 任务完成，提取结果
			result := strings.TrimPrefix(action, "finish[")
			result = strings.TrimSuffix(result, "]")
			fmt.Printf("步骤%d：思维：%s | 行动：%s | 观察：任务完成\n", step+1, thought, action)
			return result, contextManager.GetContextStr()
		}

		// 解析工具名称和参数
		executed := false
		for toolName, tool := range toolMap {
			if strings.HasPrefix(action, toolName+"[") {
				// 提取参数
				paramStart := len(toolName) + 1
				paramEnd := strings.Index(action, "]")
				if paramEnd > paramStart {
					params := action[paramStart:paramEnd]
					observation = tool.Run(params)
					executed = true
					break
				}
			}
		}

		if !executed {
			// 无效行动
			validTools := make([]string, 0, len(toolMap))
			for name := range toolMap {
				validTools = append(validTools, name)
			}
			observation = fmt.Sprintf("无效行动：%s，支持的工具为%v", action, validTools)
		}

		// 5. 更新上下文
		contextManager.AddTAO(thought, action, observation)
		fmt.Printf("步骤%d：思维：%s | 行动：%s | 观察：%s\n", step+1, thought, action, observation)
	}

	// 超时终止
	return fmt.Sprintf("任务未完成（已达最大步数%d）", maxSteps), contextManager.GetContextStr()
}

// ==================== 主程序 ====================

func main() {
	// 初始化工具
	tools := []BaseTool{
		NewFlightSearchTool(),
		NewFlightBookTool(),
	}

	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          ReAct 框架 Go 实现演示                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// ========== Good Example：成功完成任务 ==========
	fmt.Println("\n【示例 1：成功场景 - Good Example】")
	task1 := "查询明天从深圳到海南的航班，选最便宜、航班时间在晚上的那班并预订"
	result1, trajectory1 := ReactCoreLoop(task1, tools, 6)
	fmt.Printf("\n✅ 最终结果：%s\n", result1)
	fmt.Printf("\n📋 完整执行轨迹：\n%s\n", trajectory1)

	// ========== Bad Example：任务失败场景 ==========
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("\n【示例 2：失败场景 - Bad Example】")
	task2 := "查询明天从北京到火星的航班并预订"
	result2, trajectory2 := ReactCoreLoop(task2, tools, 6)
	fmt.Printf("\n❌ 最终结果：%s\n", result2)
	fmt.Printf("\n📋 完整执行轨迹：\n%s\n", trajectory2)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("演示完成！")
}
