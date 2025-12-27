# ReAct 框架 Go 实现

## 📖 项目简介

这是一个用 Go 语言实现的 **ReAct (Reasoning + Acting)** 智能体框架的演示项目。ReAct 是一种将推理（Reasoning）和行动（Acting）结合的 AI 智能体范式，通过 **"思维 → 行动 → 观察"** 的循环迭代，逐步完成复杂任务。

本项目是纯 Go 实现的教学示例，**不依赖任何外部 LLM API**，通过模拟 LLM 输出来展示 ReAct 的核心工作原理。

---

## 🎯 核心概念

### ReAct 循环（TAO Loop）

ReAct 框架的核心是 **TAO 循环**：

1. **Thought（思维）**：分析当前状态，规划下一步行动
2. **Action（行动）**：调用工具执行具体操作
3. **Observation（观察）**：获取工具执行结果，更新认知

通过不断迭代 TAO 循环，智能体能够：
- 🎯 动态调整策略
- 🔄 从错误中学习
- 🧠 处理复杂多步骤任务

---

## 🏗️ 架构设计

```
┌─────────────────────────────────────────────┐
│           ReAct Core Loop                    │
│  (TAO循环调度 + 上下文管理 + 任务控制)        │
└─────────────────┬───────────────────────────┘
                  │
         ┌────────┴────────┐
         │                 │
    ┌────▼────┐      ┌────▼────┐
    │  LLM    │      │  Tools  │
    │ 模拟器   │      │  工具层  │
    └─────────┘      └────┬────┘
                          │
              ┌───────────┴───────────┐
              │                       │
         ┌────▼─────┐          ┌─────▼────┐
         │ Flight   │          │ Flight   │
         │ Search   │          │ Book     │
         └──────────┘          └──────────┘
```

### 模块说明

#### 1. **工具封装层（Tool Layer）**
- `BaseTool` 接口：定义工具的标准化调用规范
- `FlightSearchTool`：航班查询工具
- `FlightBookTool`：航班预订工具

#### 2. **上下文管理器（Context Manager）**
- 存储 TAO 历史轨迹
- 智能裁剪超长上下文（保留近期3轮 + 早期摘要）
- 生成 LLM 可理解的上下文字符串

#### 3. **LLM 模拟器（Simulated LLM）**
- 根据任务类型和当前步骤生成预设的思维和行动
- 模拟真实 LLM 的推理过程
- 支持 Good/Bad Example 的不同响应策略

#### 4. **ReAct 核心循环（Core Loop）**
- 控制 TAO 迭代流程
- 解析 LLM 输出的思维和行动
- 调用工具并更新上下文
- 处理任务完成和超时终止

---

## 🚀 快速开始

### 环境要求
- Go 1.16 或更高版本

### 运行示例

```bash
# 克隆或下载代码
cd /path/to/react-demo

# 直接运行
go run react_demo.go

# 或编译后运行
go build -o react_demo react_demo.go
./react_demo
```

---

## 📝 示例演示

### 示例 1：成功场景（Good Example）

**任务：** 查询明天从深圳到海南的航班，选最便宜、航班时间在晚上的那班并预订

**执行流程：**

```
步骤1：
  思维：当前任务是查询明天从深圳到海南的航班，选最便宜、航班时间在晚上的那班并预订，无历史数据。
        需先调用航班查询工具，参数为"深圳,海南,明天,晚上"，获取符合条件的航班列表
  行动：flight_search[深圳,海南,明天,晚上]
  观察：符合条件航班列表：1. HU7089（深圳宝安→海口美兰，20:15-21:45，票价480元）；
                        2. CZ6753（深圳宝安→三亚凤凰，21:30-23:05，票价620元）；
                        3. MU2478（深圳宝安→海口美兰，19:40-21:10，票价550元）

步骤2：
  思维：已获取明天深圳到海南晚上的航班列表，从观察结果可知最便宜的是HU7089（票价480元）。
        下一步需调用航班预订工具，参数包含航班号HU7089、乘客信息
  行动：flight_book[HU7089,李四,440301199505056789]
  观察：航班预订成功：航班号HU7089，乘客李四（身份证号：6789），请携带有效证件提前2小时到机场办理登机手续

步骤3：
  思维：已完成最便宜晚上航班的查询与预订，观察结果显示预订成功，所有任务目标均达成，可提交最终结果
  行动：finish[明天从深圳到海南最便宜的晚上航班为HU7089（深圳宝安→海口美兰，20:15-21:45，票价480元），
              已完成预订，乘客李四（身份证号：6789）]
  观察：任务完成
```

**✅ 最终结果：** 成功完成查询和预订

---

### 示例 2：失败场景（Bad Example）

**任务：** 查询明天从北京到火星的航班并预订

**执行流程：**

```
步骤1：
  思维：需要查询北京到火星的航班信息，调用航班查询工具
  行动：flight_search[北京,火星,明天,上午]
  观察：未检索到相关航班信息（参数：北京,火星,明天,上午）

步骤2：
  思维：查询失败，尝试使用不存在的工具
  行动：rocket_search[北京,火星]
  观察：无效行动：rocket_search[北京,火星]，支持的工具为[flight_search flight_book]

步骤3：
  思维：工具调用失败，尝试直接预订
  行动：flight_book[MARS001]
  观察：航班预订失败：参数格式错误，需要'航班号,乘客姓名,身份证号'

步骤4：
  思维：多次尝试失败，任务无法完成
  行动：finish[无法完成任务：不存在北京到火星的航班]
  观察：任务完成
```

**❌ 最终结果：** 无法完成任务（不存在北京到火星的航班）

---

## 💡 核心特性

### 1. 模块化工具封装
- 统一的 `BaseTool` 接口设计
- 易于扩展新工具（只需实现 `GetName()`、`GetDescription()`、`Run()` 方法）
- 工具调用格式标准化：`tool_name[params]`

### 2. 智能上下文管理
- 自动裁剪超长上下文
- 保留近期3轮完整轨迹
- 生成早期轨迹摘要，防止信息丢失

### 3. 模拟 LLM 推理
- 根据任务类型动态生成思维和行动
- 支持多步骤任务规划
- 可轻松替换为真实 LLM API（如 OpenAI、Claude）

### 4. 鲁棒的错误处理
- 无效工具调用检测
- 参数格式校验
- 最大步数限制防止死循环

---

## 🔧 扩展开发

### 添加新工具

1. 定义工具结构体：
```go
type MyCustomTool struct {
    name        string
    description string
}
```

2. 实现 `BaseTool` 接口：
```go
func (t *MyCustomTool) GetName() string {
    return "my_tool"
}

func (t *MyCustomTool) GetDescription() string {
    return "工具功能描述"
}

func (t *MyCustomTool) Run(params string) string {
    // 实现工具逻辑
    return "执行结果"
}
```

3. 注册到工具列表：
```go
tools := []BaseTool{
    NewFlightSearchTool(),
    NewFlightBookTool(),
    NewMyCustomTool(), // 新增工具
}
```

### 接入真实 LLM

将 `SimulatedLLM` 替换为真实 LLM 调用：

```go
// 示例：调用 OpenAI API
func (llm *RealLLM) Generate(context string) string {
    prompt := buildPrompt(context) // 构建提示词
    response := callOpenAI(prompt)  // 调用 API
    return response
}
```

## 🎓 学习价值

本项目适合：
- ✅ 理解 ReAct 框架的核心工作原理
- ✅ 学习 Go 语言的接口设计和模块化编程
- ✅ 掌握智能体系统的基本架构
- ✅ 为接入真实 LLM 打下基础

---

## 📚 参考资料

- [ReAct 论文](https://arxiv.org/abs/2210.03629): "ReAct: Synergizing Reasoning and Acting in Language Models"
- [LangChain ReAct 文档](https://python.langchain.com/docs/modules/agents/agent_types/react)
- [Agent 全面爆发，一文搞懂背后的核心范式 ReAct](https://mp.weixin.qq.com/s/YQfqLoL1Z94yx9z48CE8bQ) 腾讯云开发者

