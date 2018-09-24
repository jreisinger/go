// Ex. 4.2: Print SHA256 (or SHA512) of the standard input.
package main

import (
    "io/ioutil"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
    "os"
    "flag"
    "log"
)

var width = flag.Int("w", 256, "hash width (256 or 512)")

func main() {
    flag.Parse()
    var function func(b []byte) []byte
    switch *width {
    case 256:
        function = func(b []byte) []byte {
            h := sha256.Sum256(b)
            return h[:]
        }
    case 512:
        function = func(b []byte) []byte {
            h := sha512.Sum512(b)
            return h[:]
        }
    default:
        log.Fatal("Unexpected width specified.")
    }
    b, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%x\n", function(b))
}
