// Fetch URLs in parallel and report their times and sizes.
//
// The Go Programming Language, 1.6
//
// $ go build gopl.io/ch1/fetchall
// $ ./fetchall https://golang.org http://perl.org http://reisinge.net

package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {                   // main runs in a goroutine
    start := time.Now()

    ch := make(chan string)     // channel -- a communication mechanism that
                                // allows one goroutine to pass values of a
                                // specified type to another goroutine

    for _, url := range os.Args[1:] {
        go fetch(url, ch)       // create additional goroutine
                                // goroutine -- a concurrent function execution
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)       // receive from channel ch
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)   // send to channel ch 
        return
    }

    // io.Copy function reads the body of the response and discards it by
    // writing to the ioutil.Discard output stream.
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)

    resp.Body.Close()           // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
