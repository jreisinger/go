package main

import (
	"fmt"
	"sort"
)

// This is how you sort your own data.

// The sort.Sort function takes a sort.Interface and sorts it.
// The sort.Interface requires three methods:
// * Len()
// * Less()
// * Swap()

type Person struct {
	Name string
	Age  int
}

type ByName []Person

// Return the length if the thing we are sorting.
func (ps ByName) Len() int {
	return len(ps)
}

// Is item at position i less than the item at position j?
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

// Swap the items.
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}
func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}
func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Eli", 2},
		{"Dado", 2},
		{"Matu", 4},
		{"Z", 1},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Printf("%v\n", kids)
}
