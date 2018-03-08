package main

import ( "fmt" )

func arrays() {
    var a [5]int

    a[0] = 1
    a[4] = 100

    fmt.Println(a)

    for i := 0; i < len(a); i++ {
        fmt.Println(a[i])
    }

    for i, value := range a {
        fmt.Println(i, value)
    }
}

// Like arrays, slices are indexable and have a length.
// Unlike arrays, this length is allowed to change.
func slices() {
    // make a slice (associated with an underlying array) of length 5
    s := make([]int, 10)

    s[9] = 5

    for i, value := range s {
        fmt.Println(i, value)
    }
}

func main() {
    //arrays()
    slices()
}
