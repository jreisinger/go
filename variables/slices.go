package main

import (
	"fmt"
	"os"
)

var elements = map[string]map[string]string{
	"H": map[string]string{
		"name": "Hydrogen",
		"state": "gas",
	},
	"He": map[string]string{
		"name": "Helium",
		"state": "gas",
	},
	"Li": map[string]string{
		"name": "Lithium",
		"state": "solid",
	},
}

func main() {
	for _, element := range os.Args[1:] {
		if e, ok := elements[element]; ok {
			fmt.Printf("%s is %s (%s)\n", element, e["name"], e["state"])
		} else {
			fmt.Println("I don't know", element)
		}
	}
}
