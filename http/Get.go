// Get the representation of a resource.
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "url")
		os.Exit(2)
	}
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// In fetch.go we read the body using ioutil.ReadAll()
	var buf [515]byte
	for {
		n, err := resp.Body.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
}
