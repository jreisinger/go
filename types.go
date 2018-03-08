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

    // slice can never be longer than the underlying array
    s[9] = 5

    for i, value := range s {
        fmt.Println(i, value)
    }

    // another way to create a slice
    s2 := []string{"a", "b", "c"}
    fmt.Println("Slice length:",   len(s2))
    fmt.Println("Slice capacity:", cap(s2))

    // another way to create a slice
    arr := [5]float64{1,2,3,4,5}
    s3 := arr[0:5] // [0:], [:5], [:]
    fmt.Println(s3)
}

func main() {
    //arrays()
    slices()
}
