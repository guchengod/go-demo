package main

import (
	"fmt"
	"sync"
)

func workerWithWg(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", i)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerWithWg(i, &wg)
	}
	wg.Wait()
	fmt.Println("done")
}
