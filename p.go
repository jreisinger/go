package main

import "fmt"

// returns a pointer to int (*int), i.e. an address where an int is stored
func f() *int {
    v := 1
    return &v           // address-of operator
}

func main() {
    var p = f()         // pointer
    fmt.Println(p)      // location (address)
    fmt.Println(*p)     // value (dereference)

    *p = 2              // assign new value
    fmt.Println(p)      // same address as before ^
    fmt.Println(*p)     // 2
}
