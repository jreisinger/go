package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
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
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuos, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
