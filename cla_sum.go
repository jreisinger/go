// Print the sum of command line arguments
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    args := os.Args
    sum := 0
    for i := 1; i < len(args); i++ {
        temp, _ := strconv.Atoi(args[i])
        sum += temp
    }
    fmt.Println(sum)
}
