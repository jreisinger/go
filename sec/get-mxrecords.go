package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No domain name provided.")
	}
	domain := os.Args[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, mx := range mxRecords {
		fmt.Printf("Host: %s\tPreference: %d\n", mx.Host, mx.Pref)
	}
}
