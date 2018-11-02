package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No hostname provided.")
	}
	hostname := os.Args[1]

	fmt.Println("Looking up IP address for: " + hostname)

	ips, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
