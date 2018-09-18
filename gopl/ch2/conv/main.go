// Ex. 2.2: Write a general-purpose unit-conversion program that reads numbers
// from its command-line arguments or from the standard input if there are no
// arguments, and converts each number into units like temparature in Celsius
// and Fahrenheit, length in feer and meters, weight in pounds and kilograms.

package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/jreisinger/go/gopl/ch2/tempconv"
)

func main() {
    if len(os.Args) > 1 {   // we have command-line args
        for _, arg := range os.Args[1:] {

            i, err := strconv.Atoi(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "conv: %v\n", err)
                os.Exit(1)
            }

            c := tempconv.Celsius(i)
            f := tempconv.Fahrenheit(i)

            fmt.Printf("%s = %s, %s = %s\n",
                f, tempconv.FToC(f), c, tempconv.CToF(c))
        }
    }
}
