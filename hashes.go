package main

// A hash function takes a set of data and reduces it to a smaller fixed size.

// Hash functions in Go are broken into two categories:
// * non-cryptographic (ex. hash/adler32, hash/crc32, hash/fnv)
// * cryptographic - hard to reverse (ex. crypto/sha1)

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// create a hasher
	h := crc32.NewIEEE()
	// write out data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}
