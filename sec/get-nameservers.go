package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No domain name supplied.")
	}
	domain := os.Args[1]

	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, ns := range nsRecords {
		fmt.Println(ns.Host)
	}
}
