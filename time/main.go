package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 格式化：time.Time -> string
	layout := "2006-01-02 15:04:05"
	formatted := now.Format(layout)
	fmt.Println("Formatted:", formatted)

	// 解析：string -> time.Time
	timeStr := "2025-10-13 20:30:00"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed:", parsedTime)

	// 10秒后
	tenSecondsLater := now.Add(10 * time.Second)
	// 1小时前
	oneHourAgo := now.Add(-1 * time.Hour)

	// 计算两个时间的差值
	diff := tenSecondsLater.Sub(oneHourAgo) // 返回一个 Duration
	fmt.Printf("Difference is %.2f hours\n", diff.Hours())

	// 睡眠（阻塞当前 goroutine）
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Awake!")

	fmt.Println("Starting a one-shot timer for 2 seconds...")
	timer := time.NewTimer(2 * time.Second)

	// 在另一个 goroutine 中等待...
	go func() {
		<-timer.C // 这里会被阻塞，直到定时器触发
		fmt.Println("Timer fired!")
	}()

	// 主 goroutine 可以继续做其他事情...
	time.Sleep(3 * time.Second) // 等待演示结束

	// 创建一个每 500 毫秒触发一次的 ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop() // 用完后别忘了停止，以释放资源

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Ticker stopped.")
				return
			case t := <-ticker.C: // 每 500ms，这个 case 会被触发
				fmt.Println("Tick at", t.Format("15:04:05.000"))
			}
		}
	}()

	// 让 ticker 运行 2 秒
	time.Sleep(2 * time.Second)
	done <- true                      // 发送停止信号
	time.Sleep(50 * time.Millisecond) // 等待 goroutine 退出

}
