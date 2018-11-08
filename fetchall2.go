// This is from video course Go Introduction (I think).
// I named it fetchall2.go since it's similar to fetchall.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// getPageSize returns the size in bytes of a page found at url.
func getPageSize(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, err
	}

	return len(body), err
}

// worker is a wrapper for getPageSize to enable concurrency.
func worker(urlCh chan string, msg chan string, id int) {
	url := <-urlCh
	length, err := getPageSize(url)
	if err == nil {
		msg <- fmt.Sprintf("%s has %d bytes (worker %d)\n", url, length, id)
	} else {
		msg <- fmt.Sprintf("Error getting %s: %s (worker %d)\n", url, err, id)
	}
}

func main() {
	urlCh := make(chan string)
	msgCh := make(chan string)

	// Start workers
	for i := 0; i < len(os.Args)-1; i++ {
		fmt.Fprintf(os.Stderr, "Started worker %d\n", i)
		go worker(urlCh, msgCh, i)
	}

	// Send urls to workers
	for _, url := range os.Args[1:] {
		urlCh <- url
	}

	// Get results from workers
	for range os.Args[1:] {
		fmt.Printf("%s", <-msgCh)
	}
}
