package main

import "fmt"

func squareAddress(a *int) *int {
	*a *= *a
	fmt.Println("\nInside squareAddress")
	fmt.Printf("Value of a is: %d \n", a)
	fmt.Printf("Value of &a is: %d \n", &a) // address of pointer variable
	fmt.Printf("Value of *a is: %d \n", *a)
	return a
}
func main() {

	a := 3

	fmt.Println("\nInside main before squareAddress")
	fmt.Printf("Value of a is: %d \n", a)
	fmt.Printf("Value of &a is: %d \n", &a)
	// fmt.Printf("Value of *a is: %d \n", *a)

	// Reference of the value - pass by reference
	var reference = squareAddress(&a)

	fmt.Println("\nInside main after squareAddress")
	fmt.Printf("Value of a is: %d \n", a)
	fmt.Printf("Value of &a is: %d \n", &a)
	// fmt.Printf("Value of *a is: %d \n", *a)
	fmt.Printf("Value of reference is: %d \n", reference)
	fmt.Printf("Value of *reference is: %d \n", *reference)
}
