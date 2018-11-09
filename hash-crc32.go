// A hash function takes a set of data and reduces it to a smaller fixed size.

// Hash functions in Go are broken into two categories:
// * non-cryptographic (ex. hash/adler32, hash/crc32, hash/fnv)
// * cryptographic - hard to reverse (ex. crypto/sha1)

// A common use for crc32 is to compare two files. If the Sum32 value for both
// files is the same, it's highly likely that the file are the same.

// Taken from Introducing Go.

package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func getHash(fileName string) (uint32, error) {
	// open the file
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// create a hasher
	h := crc32.NewIEEE()
	// copy the file into the hasher
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {
	file1 := os.Args[1]
	file2 := os.Args[2]

	h1, err := getHash(file1)
	if err != nil {
		log.Fatal(err)
	}

	h2, err := getHash(file2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h1, h2, h1 == h2)
}
