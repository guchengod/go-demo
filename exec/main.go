package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	//Cmd 结构体有很多有用的字段可以配置，最常用的有：
	//
	//cmd.Path: 命令的可执行文件路径（exec.Command 通常会帮你自动在 PATH 环境变量中找到）。
	//
	//cmd.Args: 传递给命令的完整参数列表，包括命令本身作为第一个元素。
	//
	//cmd.Stdout, cmd.Stderr: io.Writer 类型，可以用来指定命令的标准输出和标准错误输出到哪里（比如 os.Stdout, os.Stderr, 或者一个文件）。
	//
	//cmd.Stdin: io.Reader 类型，用来向命令提供标准输入。
	//
	//cmd.Dir: string 类型，指定命令的工作目录。
	//等待命令完成
	// 准备一个命令
	cmd := exec.Command("sleep", "2")
	fmt.Println("Running 'sleep 2'...")

	// Run 会启动命令并阻塞，直到命令执行完成
	err := cmd.Run()

	// 错误处理非常重要！
	// 如果 err != nil，有两种可能：
	// 1. 命令无法启动（比如 "sleeeep" 命令不存在）
	// 2. 命令启动了，但以非 0 的状态码退出（表示执行失败）
	if err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
		// 如果是退出状态码错误，可以这样获取更具体的信息
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exit code is %d\n", exitError.ExitCode())
		}
		return
	}

	fmt.Println("Command finished successfully.")

	// 获取输出
	// 准备命令 `git rev-parse --short HEAD` 来获取当前 git commit 的短哈希
	cmd = exec.Command("git", "rev-parse", "--short", "HEAD")

	// Output 会运行命令，并返回一个包含标准输出的 []byte
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		// 当使用 Output 时，如果命令失败，err (*exec.ExitError) 会包含 stderr 的输出
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Stderr: %s\n", string(exitError.Stderr))
		}
		return
	}

	// output 是 []byte，通常需要转换为 string 并去掉末尾的换行符
	commitHash := strings.TrimSpace(string(output))
	fmt.Printf("Current commit hash: %s\n", commitHash)

	// 如果你想同时捕获 stdout 和 stderr，可以使用 cmd.CombinedOutput()

	//流式处理与完全控制 (Start, Wait, Pipes)
	// 准备一个会持续输出的命令
	cmd = exec.Command("ping", "-c", "4", "baidu.com")

	// 1. 获取命令的标准输出管道
	// 必须在 Start() 之前调用
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// 2. 启动命令（Start 是非阻塞的！）
	// 它会立即返回，命令在后台开始执行
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Command started in background...")

	// 使用 bufio.Scanner 来实时读取输出流
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// 每当命令产生一行新的输出，这里就会打印出来
		fmt.Println("Ping output:", scanner.Text())
	}

	// 3. 等待命令执行完成并清理资源
	// Wait() 会阻塞，直到命令结束
	// 必须调用 Wait()，否则可能会产生僵尸进程
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait returned an error:", err)
	}
	fmt.Println("Command finished.")
}
