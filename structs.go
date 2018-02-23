package main

import ("fmt"; "math")

// New type definition
//// type   - defines new type
//// Circle - name of the type
//// struct - type that contains named fields
type Circle struct {
    //x float64
    //y float64
    //r float64
    x, y, r float64
}

// Function using the Circle
func circleArea(c Circle) float64 {
    return math.Pi * c.r*c.r
}

func main() {

    // Initialization of a variable of type Circle
    //var c Circle        // fields set to zeros
    //c := new(Circle)    // returns pointer to the struct
    c := &Circle{1,2,3}

    // Accessing the struct fields
    c.r = 5

    // Using the struct as a function argument
    fmt.Println(circleArea(*c))
}

