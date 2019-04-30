package main

import (
	"fmt"
	"math"
)

// Create a few small structs that have what you want, add in methods that you
// need, and as you build your program, useful interfaces will tend to emerge.
// There's no need to have them all figured out ahead of time.

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Interface defines a method set, i.e. the list of methods 
// a type must have in order to implement the interface.
type Shape interface {
	area() float64
	//perimeter() float64
}

// Interface type used here as an argument to a function.
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 1}
	fmt.Printf("Area of cirle %v: %f\n", c, c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Printf("Aread of rectangle %v: %f\n", r, r.area())

	fmt.Printf("Total area: %f\n", totalArea(&c, &r))
}
