package main

import "fmt"

// Escape Analysis: How to allocate memory in stack & heap

// stack allocation
func add(a, b int) int {
	return a + b
}

func test() *int {
	number := 10
	// number escape to heap
	return &number
}

func main() {

	fmt.Printf("Addition: %d \n", add(10, 20))
	var ptr *int = test()
	fmt.Printf("Value of ptr: %d \n", ptr)
	fmt.Printf("Value of *ptr: %d \n", *ptr)
}
