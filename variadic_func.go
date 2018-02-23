package main

import "fmt"

// args is a variadic parameter
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1,2))
    fmt.Println(add(1,2,3))
}
