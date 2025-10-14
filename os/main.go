package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//readSmallFile()
	readLargeFile()
}

func readLargeFile() {
	open, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func readSmallFile() {

	file, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(file))

}
