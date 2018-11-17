package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Android struct {
	Person // embedded type (or anonymous field)
	Model  string
}

// Method is just a function with a receiver (old OO lingo).
func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	// Object instantiation - way 1
	p := Person{Name: "John"}

	// Object instantiation - way 2
	a := new(Android)
	a.Name = "R2D2"

	// A method call
	p.talk()
	a.talk()
}
