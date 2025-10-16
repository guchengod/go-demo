package main

import (
	"flag"
	"fmt"
)

func main() {
	// --- 1. 定义标志 (我们使用方法A) ---
	port := flag.Int("port", 8080, "The port to listen on.")
	host := flag.String("host", "localhost", "The host to connect to.")
	verbose := flag.Bool("verbose", false, "Enable verbose output.")

	// --- 2. 解析标志 ---
	// 必须在所有 flag 定义之后，并且在使用 flag 之前调用
	flag.Parse()

	// --- 3. 使用标志的值 ---
	// 注意：因为 port, host, verbose 是指针，所以需要用 * 解引用
	fmt.Println("--- Flag Values ---")
	fmt.Printf("Port: %d\n", *port)
	fmt.Printf("Host: %s\n", *host)
	fmt.Printf("Verbose Mode: %t\n", *verbose)

	// 获取并打印所有非标志的普通参数
	otherArgs := flag.Args()
	fmt.Println("\n--- Other Arguments ---")
	if len(otherArgs) > 0 {
		for i, arg := range otherArgs {
			fmt.Printf("Arg %d: %s\n", i, arg)
		}
	} else {
		fmt.Println("No other arguments provided.")
	}
}
