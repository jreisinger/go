package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print device information
	for _, device := range devices {
		fmt.Printf("%s (%s)\n", device.Name, device.Description)
		for _, addr := range device.Addresses {
			fmt.Printf("* %s/%d\n", addr.IP, addr.Netmask)
		}
	}
}
