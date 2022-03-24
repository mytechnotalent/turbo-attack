// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides numeric conversion functionality.
package convert

import (
	"encoding/binary"
	"net"
	"strconv"
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
func IPV4(ip, port, count *string) (*[]byte, *[]byte, *int) {
	ipv4NetIP := net.ParseIP(*ip)
	ipv4Byte := []byte(ipv4NetIP)
	portInt, err := strconv.Atoi(*port) // convert first string octet to int
	if err != nil {
		panic(err)
	}
	portByte := IntToByte(portInt)
	countInt, err := strconv.Atoi(*count)
	if err != nil {
		panic(err)
	}
	return &ipv4Byte, &portByte, &countInt
}

// IPV6 converts a string into a properly formatted byte and int.
// It will return a byte and int.
func IPV6(ip, port, count *string) (*[]byte, *[]byte, *int) {
	ipv6NetIP := net.ParseIP(*ip)
	ipv6Byte := []byte(ipv6NetIP)
	portInt, err := strconv.Atoi(*port) // convert first string octet to int
	if err != nil {
		panic(err)
	}
	portByte := IntToByte(portInt)
	countInt, err := strconv.Atoi(*count)
	if err != nil {
		panic(err)
	}
	return &ipv6Byte, &portByte, &countInt
}
