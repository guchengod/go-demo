// file: main.go
package main

import (
	"fmt"
	"regexp"
)

// 1. 在包级别编译一次正则表达式
// 定义捕获组：
// Group 1: (.*?)  - 匹配时间戳
// Group 2: (\w+)   - 匹配日志级别 (ERROR, INFO, etc.)
// Group 3: (.*)   - 匹配剩余的所有消息
var logParser = regexp.MustCompile(`\[(.*?)\] \[(\w+)\] (.*)`)

func parseLogLine(line string) {
	fmt.Printf("--- Parsing line: \"%s\" ---\n", line)

	// 2. 使用编译好的对象进行匹配和提取
	match := logParser.FindStringSubmatch(line)

	// 3. 检查是否匹配成功并使用结果
	if match == nil {
		fmt.Println("Line does not match the pattern.")
		return
	}

	// match[0] 是整个匹配的字符串
	// match[1], match[2], match[3] 分别对应三个捕获组
	timestamp := match[1]
	level := match[2]
	message := match[3]

	fmt.Printf("Timestamp: %s\n", timestamp)
	fmt.Printf("Level:     %s\n", level)
	fmt.Printf("Message:   %s\n", message)
}

func main() {
	line1 := "[2025-10-13 20:53:48] [ERROR] User login failed: invalid password"
	line2 := "[2025-10-13 20:54:01] [INFO] Server started on port 8080"
	line3 := "This is not a valid log line."

	parseLogLine(line1)
	fmt.Println()
	parseLogLine(line2)
	fmt.Println()
	parseLogLine(line3)
}
