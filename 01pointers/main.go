package main

import "fmt"

func main() {

	// Initial
	var a int = 10

	var ptr *int = &a

	fmt.Println("Initial values")
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Value of &a: %d\n", &a)
	fmt.Printf("value of ptr: %d \n", ptr)
	fmt.Printf("value of *ptr: %d \n", *ptr)

	// Re-initialization
	a = 20
	fmt.Println("After reinitialization values")
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Value of &a: %d\n", &a)
	fmt.Printf("value of ptr: %d \n", ptr)
	fmt.Printf("value of *ptr: %d \n", *ptr)

	//Change the value of x via pointer: Dereferencing
	*ptr = 30
	fmt.Println("After Dereferencing values")
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Value of &a: %d\n", &a)
	fmt.Printf("value of ptr: %d \n", ptr)
	fmt.Printf("value of *ptr: %d \n", *ptr)

}
