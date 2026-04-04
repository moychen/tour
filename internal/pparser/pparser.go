package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// func main() {
// 	err := cmd.Execute()
// 	if err != nil {
// 		log.Fatalf("Cmd.Execute err: %v", err)
// 	}
// }

// Use tcpdump to create a test file
// tcpdump -w test.pcap
// or use the example above for writing pcap files

var (
	pcapFile string = "../../../408.pcapng"
	handle   *pcap.Handle
	err      error
)

func main() {
	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		/*
			ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
			if ethernetLayer != nil {
				fmt.Println("Ethernet layer detected.")
				ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
				fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
				fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
				// Ethernet type is typically IPv4 but could be ARP or other
				fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
				fmt.Println("111111111111")
			}
		*/

		/*
			// Let's see if the packet is IP (even though the ether type told us)
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				fmt.Println("IPv4 layer detected.")
				ip, _ := ipLayer.(*layers.IPv4)

				// IP layer variables:
				// Version (Either 4 or 6)
				// IHL (IP Header Length in 32-bit words)
				// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
				// Checksum, SrcIP, DstIP
				fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
				fmt.Println("Protocol: ", ip.Protocol)
				fmt.Println("1111111111111")
			}
		*/

		// Let's see if the packet is TCP
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)

			// TCP layer variables:
			// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
			// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS

			if (tcp.SrcPort == 19031 && tcp.DstPort == 7791) || (tcp.DstPort == 19031 && tcp.SrcPort == 7791) {
				fmt.Println("TCP layer detected.")
				fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
				fmt.Println("Sequence number: ", tcp.Seq)
				fmt.Println(string(tcpLayer.LayerPayload()))
				fmt.Println("11111111111")

				applicationLayer := packet.ApplicationLayer()
				if applicationLayer != nil {
					fmt.Println(string(applicationLayer.Payload()))
					fmt.Println("22222222222")
				}
			}
		}
	}

	// var eth layers.Ethernet
	// var ip4 layers.IPv4
	// var tcp layers.TCP

	// parser := gopacket.NewDecodingLayerParser(
	// 	layers.LayerTypeEthernet, &eth, &ip4, &tcp)
	// decodedLayers := []gopacket.LayerType{}

	// 解析
	// for packet := range packetSource.Packets() {
	// 	parser.DecodeLayers(packet, &decodedLayers)
	// 	for _, layerType := range decodedLayers {
	// 		fmt.Println(layerType)
	// 	}
	// }

}
