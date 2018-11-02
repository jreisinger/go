package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"os"
	"time"
)

var (
	device            = "enp0s31f6"
	snapshotLen int32 = 1024
	promiscuos        = false
	err         error
	timeout     = -1 * time.Second
	handle      *pcap.Handle
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuos, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	if len(os.Args) > 1 {
		filter := os.Args[1]
		err := handle.SetBPFFilter(filter)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get and print packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
