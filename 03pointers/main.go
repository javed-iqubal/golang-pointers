package main

import "fmt"

func squareValue(a int) int {
	a *= a
	fmt.Println("\nInside squareValue")
	fmt.Printf("Value of a is: %d \n", a)
	fmt.Printf("Address of a is: %d \n", &a)
	return a
}
func main() {

	a := 3
	// copy of the value - pass by value
	fmt.Println("\nInside main before squareValue")
	fmt.Printf("Value of a is: %d \n", a)
	fmt.Printf("Address of a is: %d \n", &a)
	squareValue(a)
	fmt.Println("\nInside main after squareValue")
	fmt.Printf("Value of a is: %d \n", a)
	fmt.Printf("Address of a is: %d \n", &a)
}
