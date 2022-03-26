// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides input validation and conversion functionality.
package convert

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"strings"
)

// IntToByte converts an int to a byte.
// It will return a byte.
func IntToByte(n int) []byte {
	nUint64 := uint64(n)
	nByte := make([]byte, 8)
	binary.BigEndian.PutUint64(nByte, nUint64)
	return nByte
}

// IPV4 converts a string into a properly formatted byte and int.
// It will return a byte and int.
func IPV4(ethInterface, ip, port, count *string) (*[]byte, *[]byte, *int) {
	if !((*ethInterface == "eth0") || (*ethInterface == "wlan0")) {
		log.Fatal("invalid interface")
	}
	ipv4NetIP := net.ParseIP(*ip)
	if ipv4NetIP == nil {
		log.Fatal("invalid ip4 ip")
	}
	validIP4IP := strings.Contains(*ip, ":")
	if validIP4IP == true {
		log.Fatal("invalid ip4 ip")
	}
	ipv4Byte := []byte(ipv4NetIP)
	portInt, err := strconv.Atoi(*port) // convert first string octet to int
	if err != nil {
		log.Fatal("invalid port")
	}
	portByte := IntToByte(portInt)
	countInt, err := strconv.Atoi(*count)
	if err != nil {
		log.Fatal("invalid count")
	}
	return &ipv4Byte, &portByte, &countInt
}

// IPV6 converts a string into a properly formatted byte and int.
// It will return a byte and int.
func IPV6(ethInterface, ip, port, count *string) (*[]byte, *[]byte, *int) {
	if !((*ethInterface == "eth0") || (*ethInterface == "wlan0")) {
		log.Fatal("invalid interface")
	}
	ipv6NetIP := net.ParseIP(*ip)
	if ipv6NetIP == nil {
		log.Fatal("invalid ip6 ip")
	}
	validIP6IP := strings.Contains(*ip, ".")
	if validIP6IP == true {
		log.Fatal("invalid ip6 ip")
	}
	ipv6Byte := []byte(ipv6NetIP)
	portInt, err := strconv.Atoi(*port) // convert first string octet to int
	if err != nil {
		log.Fatal("invalid port")
	}
	portByte := IntToByte(portInt)
	countInt, err := strconv.Atoi(*count)
	if err != nil {
		log.Fatal("invalid count")
	}
	return &ipv6Byte, &portByte, &countInt
}
