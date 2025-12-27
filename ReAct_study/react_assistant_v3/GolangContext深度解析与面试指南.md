你好！我是你的 Go 语言架构专家。关于 `context`，我们不废话，直接进入核心。

### **核心结论 (TL;DR)**
`context` 是 Go 并发编程的**“任务指挥棒”**。它通过**树状传播机制**，在 Goroutine 之间传递**取消信号、超时限期和请求链路数据**，核心目的是**防止资源泄漏**（如孤儿 Goroutine）并实现**跨层级的链路追踪**。

---

### **思维导图 (ASCII Art)**

```text
      [Root: Background/TODO]
               |
      --------------------
      |        |         |
[Cancel]   [Timer]    [Value]  <-- 衍生节点 (Derivatives)
  |          |          |
信号传播    超时控制    元数据传递  <-- 核心能力 (Capabilities)
  |          |          |
递归关闭    计时触发    向上查找    <-- 运作机制 (Mechanisms)
```

---

### **深入分析 (Deep Dive)**

#### **1. 场景与类比 (费曼学习)**
*   **类比**：想象一场“接力比赛”。每个运动员（Goroutine）手里都拿着一个“对讲机”（Context）。
    *   如果教练（主协程）在对讲机里喊“比赛取消”，所有正在跑的运动员都会收到信号并立刻停下。
    *   对讲机里还贴着一张纸条（Value），记录着选手的编号（TraceID）。
*   **核心场景**：
    1.  **超时控制**：调用数据库或第三方 API 时，设置 500ms 超时，防止接口挂起导致雪崩。
    2.  **级联取消**：用户关闭了网页，后端立即停止所有为该用户执行的复杂计算和 DB 查询。
    3.  **链路追踪**：在整个微服务调用链中透传 `RequestID` 或 `AuthToken`。

#### **2. 实现哲学 (核心)**
*   **信号机制**：`context` 的灵魂是 `Done()` 方法返回的一个 **只读 Channel**。当 `cancel()` 被调用或超时发生时，该 Channel 被 `close`。利用 `close` 广播机制，所有监听该 Channel 的 Goroutine 都能瞬时收到通知。
*   **树状传播**：父 Context 取消时，会递归调用所有子 Context 的取消函数（`propagateCancel`），确保整棵任务树被清理。
*   **KV 存储**：`valueCtx` 采用简单的链表结构。查找时从当前节点**向上递归**查找父节点，直到根部。复杂度 $O(n)$，因此严禁存储大量数据。

#### **3. 源码寻踪 (精准)**
*   **文件**：`src/context/context.go`
*   **核心结构**：
    *   `cancelCtx`：实现取消功能的核心，维护 `children` map。
    *   `timerCtx`：在 `cancelCtx` 基础上多了个 `time.Timer`。
*   **核心函数**：`propagateCancel`。它负责挂载子节点到父节点上，确保取消动作的级联。

#### **4. 面试高频：伪代码模拟核心逻辑**
```go
// 监听取消的核心范式
func doWork(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            // 收到信号，清理并退出，防止内存泄漏
            return
        default:
            // 执行业务逻辑
        }
    }
}
```

#### **5. 可运行的 Demo (超时控制)**
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. 创建一个 1秒超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // 必须调用 cancel 以释放计时器资源

	// 2. 模拟耗时操作
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		// 3. 1秒后 ctx.Done() 会关闭，走到这里
		fmt.Println("任务超时:", ctx.Err()) 
	}
}
```

#### **6. 面试避坑指南 (Anti-patterns)**
*   **别存结构体**：不要把 Context 放在结构体里，它应该作为函数的第一参数显式传递。
*   **Key 冲突**：`WithValue` 的 Key 不要用字符串，应使用自定义的 `unexported` 类型，防止被其他包覆盖。
*   **必须 Cancel**：只要用了 `WithCancel` 或 `WithTimeout`，必须 `defer cancel()`，否则在 Context 过期前，父节点会一直持有子节点的引用，导致内存缓慢泄漏。
*   **Go 1.20+ 新特性**：面试时提到 `context.WithCancelCause` 可以传递具体的取消错误原因，会显得你紧跟前沿。

---

### **推荐阅读**
1.  [Go 官方文档 - Package context](https://pkg.go.dev/context)
2.  [Go Blog - Go Concurrency Patterns: Context](https://go.dev/blog/context)