// Print the command line arguments
package main

import "fmt"
import "os"

func main() {
    args := os.Args
    for i := 0; i < len(args); i++ {
        fmt.Println(args[i])
    }
}
