package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}
	fmt.Printf("len %d cap %d \n", len(letters), cap(letters))
	slice1 := letters[1:2]
	fmt.Printf("len %d cap %d \n", len(slice1), cap(slice1))
	slice1[0] = "X"
	slice2 := letters[0:2]
	fmt.Printf("len %d cap %d \n", len(slice2), cap(slice2))

	slice1 = append(slice1, "Y")
	fmt.Printf("len %d cap %d \n", len(slice1), cap(slice1))
	fmt.Printf("slice1 %v \n", slice2)
	slice1 = append(slice1, "Z", "P", "Q")
	fmt.Printf("len %d cap %d \n", len(slice1), cap(slice1))
	fmt.Printf("slice1 %v \n", slice1)
	fmt.Printf("slice2 %v \n", slice2)
}
