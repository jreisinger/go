package main

// go run -v myfile.go
//  go              -- command
//  run, myfile.go  -- arguments
//  -v              -- flag

import("fmt";"flag";"math/rand")

func main() {
    // define flags
    maxp := flag.Int("max", 6, "the max value")

    // parse
    flag.Parse()

    // any additional non-flag args can be retrieved 
    // with flag.Args(), which returns a []string

    // generate a random number between 0 and max
    fmt.Println(rand.Intn(*maxp))
}
