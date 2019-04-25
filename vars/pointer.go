package main

import "fmt"

// returns a pointer to int (*int), i.e. an address where an int is stored
func f() *int {
	v := 0    // declare and initialize a variable
	v = 1     // update the variable directly (through a name)
	return &v // address-of operator
}

func main() {
	var p = f()     // pointer
	fmt.Println(p)  // location (address)
	fmt.Println(*p) // value (dereference)

	*p = 2          // update the variable indirectly (through an address)
	fmt.Println(p)  // same address as before ^
	fmt.Println(*p) // "2"

	// each call to f() returns a distinct value
	fmt.Println(f() == f()) // "false"
}
