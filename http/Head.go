// Get HTTP headers, i.e. information about an HTTP resource as opposed to the
// representation of a resource you would get with Get().
//
// From Network Programming with Go, ch. 8
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
		os.Exit(2)
	}
	url := os.Args[1]

	resp, err := http.Head(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", resp.Status)
	for k, v := range resp.Header {
		fmt.Printf("%s %s\n", k, v)
	}
}
