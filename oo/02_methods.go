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

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	p := Person{Name: "John"}

	a := new(Android)
	a.Name = "R2D2"

    // A method call
	p.talk()
	a.talk()
}
