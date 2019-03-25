package main

import (
	"fmt"
	"os"
)

// ints is a variadic parameter
func addInts(ints ...int) int {
	total := 0
	for _, v := range ints {
		total += v
	}
	return total
}

func printArgs(args ...string) {
	for _, a := range args {
		fmt.Println(a)
	}
}

func main() {
	fmt.Println(addInts(1, 2))
	fmt.Println(addInts(1, 2, 3))

	// you can't leave out '...'
	printArgs(os.Args[1:]...)
}
