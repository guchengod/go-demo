package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) (string, error) {
	resultChan := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		resultChan <- "here is your data"
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-resultChan:
		return result, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data, err := fetchData(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		// 因为我们返回了原始的 ctx.Err()，所以调用者可以这样做：
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("This was a timeout error. Maybe we should try again later with a longer timeout.")
		} else if errors.Is(err, context.Canceled) {
			fmt.Println("This was a manual cancellation.")
		}
	} else {
		fmt.Println("Data received:", data)
	}
}
