package main

import "fmt"

// get the memory address using &
func getAddressUsingAnd() *int {
	x := 20
	return &x
}

// get the memory address using new
func getAddressUsingNew() *int {
	x := 20
	return new(x)
}

func main() {

	fmt.Printf("Using & -> address: %d, Value: %d: \n", getAddressUsingAnd(), *getAddressUsingAnd())

	fmt.Printf("Using new -> address: %d, Value: %d: \n", getAddressUsingNew(), *getAddressUsingNew())

}
