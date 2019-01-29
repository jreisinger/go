// Base64 is one common format for encoding binary data as strings.
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	eighBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	enc := base64.StdEncoding.EncodeToString(eighBitData)
	dec, _ := base64.StdEncoding.DecodeString(enc)

	fmt.Printf("Original data:\t%v\n", eighBitData)
	fmt.Printf("Encoded string:\t%v\n", enc)
	fmt.Printf("Decoded string:\t%v\n", dec)
}
