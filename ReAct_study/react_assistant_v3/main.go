package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Run 运行由代码控制的、纯 LLM 驱动的三阶段专家流程
func Run(task string, llm LLM) (string, error) {
	fmt.Printf("\n========================================\n")
	fmt.Printf("任务：%s\n", task)
	fmt.Printf("========================================\n\n")

	// ----------- 阶段 1: 信息规划师 -> 生成关键词 -----------
	fmt.Println("--- 阶段 1: 信息规划师 ---")
	fmt.Println("思考：为了深入回答这个问题，我需要从哪些角度和技术点进行分析？")
	prompt1 := fmt.Sprintf("针对问题：'%s'，请你作为一名资深技术专家，列出需要调研的核心技术关键词、关键数据结构和相关项目/模块。请只返回用逗号分隔的关键词列表，不要任何其他解释。", task)

	keywords, err := llm.Generate(prompt1)
	if err != nil {
		return "", fmt.Errorf("阶段 1 (信息规划师) LLM 调用失败: %w", err)
	}
	keywords = strings.Trim(keywords, "\n \"`")
	fmt.Printf("行动：规划出以下核心关键词 -> [ %s ]\n\n", keywords)

	// ----------- 阶段 2: 知识提炼师 -> 生成研究报告 -----------
	fmt.Println("--- 阶段 2: 知识提炼师 ---")
	fmt.Println("思考：基于这些关键词，我脑中的知识库能整理出一份怎样的核心知识摘要？")
	prompt2 := fmt.Sprintf(
		"你是一名知识渊博的研究员。请基于以下核心关键词：'%s'，并结合你对问题 '%s' 的理解，在你的知识库中进行深入分析和整理，生成一份关于这个主题的、包含核心要点和技术细节的“研究报告”。报告需要准确、深入，为最终回答做准备。", keywords, task)

	researchReport, err := llm.Generate(prompt2)
	if err != nil {
		return "", fmt.Errorf("阶段 2 (知识提炼师) LLM 调用失败: %w", err)
	}
	fmt.Printf("行动：生成了核心技术研究报告。 (内容省略)\n\n")

	// ----------- 阶段 3: 架构师 -> 生成最终答案 -----------
	fmt.Println("--- 阶段 3: 架构师与表达者 ---")
	fmt.Println("思考：现在我拥有了问题的核心关键词和一份深度研究报告，可以开始撰写最终的专家级回答了。")
	finalPrompt := fmt.Sprintf(
		"你是一位拥有超过10年经验的 Go 语言架构师和源码分析专家，并且是一位沟通大师，擅长使用金字塔原理，用最精炼的语言解释最核心的问题。\n\n"+
			"你的任务是深入、清晰、准确地回答以下问题：'%s'。\n\n"+
			"作为背景知识，你已经整理出了一份核心研究报告：\n<研究报告>\n%s\n</研究报告>\n\n"+
			"现在，请你综合以上所有信息，产出最终的、高质量的回答。"+
			"你的回答必须严格遵循金字塔原理（结论先行），并使用 Markdown 格式，所有部分都必须**高度精炼、直击核心**。\n\n"+
			"**核心结论 (TL;DR)**\n"+
			"在这里用一两句最精炼、最直白的话直接给出问题的核心答案。\n\n"+
			"**思维导图 (ASCII Art)**\n"+
			"在这里用 ASCII 字符绘制一个**简化版**的思维导图，只包含最重要的1-3个核心概念及其关系。\n\n"+
			"--- \n\n"+
			"**深入分析 (Deep Dive)**\n"+
			"在这里展开论证，但每一部分都必须非常简短、切中要害。\n"+
			"### 1. 场景与类比 (费曼学习)\n"+
			"用一个生动的、非技术的比喻（如洋葱、流水线）来解释核心概念。并列举2-3个这个技术在真实世界中的主要应用场景（如：日志、鉴权、链路追踪等）。\n\n"+
			"### 2. 实现哲学 (核心)\n"+
			"只聚焦于**最核心的1-2个数据结构和关键机制**，用简短的话解释它们如何协作。\n\n"+
			"### 3. 伪代码 (核心)\n"+
			"如果逻辑简单，可以直接用文字描述。如果必须，只展示**最关键的1个函数**的核心逻辑伪代码。\n\n"+
			"### 4. 源码寻踪 (精准)\n"+
			"只列出**最重要的1-2个**文件和函数名。\n\n"+
			"### 5. 可运行的Demo (实践)\n"+
			"提供一个完整的、可以直接运行的 Go 代码片段，用于演示这个核心概念。代码需要包含 main 函数和必要的 import。\n\n"+
			"### 6. 推荐阅读\n"+
			"提供1-2个你认为最高质量的外部链接（博客或官方文档）。", task, researchReport)

	finalAnswer, err := llm.Generate(finalPrompt)
	if err != nil {
		return "", fmt.Errorf("阶段 3 (架构师) LLM 调用失败: %w", err)
	}

	fmt.Println("✅ 任务完成")
	return finalAnswer, nil
}


// ==================== 主程序 ====================

func main() {
	// 1. 使用 flag 包从命令行接收用户问题
	var userQuestion string
	flag.StringVar(&userQuestion, "q", "", "您想深入分析的 Go 源码问题")
	flag.Parse()

	if userQuestion == "" {
		fmt.Println("用法: go run . -q \"你的问题\"")
		fmt.Println("例如: go run . -q \"请深入讲解一下 Gin 框架的中间件是如何实现的\"")
		os.Exit(1)
	}

	// 初始化 LLM
	llm := NewGoogleLLM()

	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          Go 源码架构师 (三阶段专家模式)                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// 2. 将用户问题作为任务启动新的运行流程
	result, err := Run(userQuestion, llm)
	if err != nil {
		fmt.Printf("\n\n❌ 执行过程中发生错误: %v\n", err)
		os.Exit(1)
	}

	// 3. 打印最终结果的预览
	fmt.Printf("\n\n================ 最终回答 (预览) ==================\n")
	if utf8.RuneCountInString(result) > 500 {
		// 使用 Rune 来处理多字节字符，确保不会截断汉字
		preview := string([]rune(result)[:500])
		fmt.Println(preview, "\n\n... (完整报告已保存到 .md 文件)")
	} else {
		fmt.Println(result)
	}
	fmt.Printf("====================================================\n\n")

	// 4. 为报告生成一个智能文件名
	fmt.Println("--- 行动：生成报告文件名 ---")
	fileNamePrompt := fmt.Sprintf("请根据这个问题：'%s'，生成一个简短、专业、适合用于 Markdown 文件名的标题。例如，对于 '请深入讲解一下 Gin 框架的中间件是如何实现的'，一个好的标题是 'Gin框架中间件原理核心说明'。请只返回标题文字，不要包含 '.md' 后缀或任何其他解释。", userQuestion)
	fileName, err := llm.Generate(fileNamePrompt)
	if err != nil {
		fmt.Printf("⚠️  生成智能文件名失败: %v。将使用默认文件名。\n", err)
		fileName = "report" // 回退到默认文件名
	}
	fileName = strings.Trim(fileName, "\n \"`")


	// 5. 使用 SaveToFileTool 将结果保存为 Markdown 文件
	fmt.Println("--- 行动：保存报告到文件 ---")
	saveTool := NewSaveToFileTool()
	params := fmt.Sprintf("%s|||%s", fileName, result)

	_, err = saveTool.Run(params)
	if err != nil {
		fmt.Printf("❌ 保存文件时发生错误: %v\n", err)
	}
}
