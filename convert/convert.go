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
	"errors"
	"net"
	"strconv"
	"strings"
)

// IntToByte converts an int to a byte.
// It will return a []byte or an error if one occurred.
func IntToByte(n int) ([]byte, error) {
	if n < 0 || n > 1000000 {
		return nil, errors.New("invalid int range")
	}
	nUint64 := uint64(n)
	nByte := make([]byte, 8)
	binary.BigEndian.PutUint64(nByte, nUint64)
	return nByte, nil
}

// IPV4 converts a string into a properly formatted byte and int.
// It will return a []byte and int or an error if one occurred.
func IPV4(ethInterface, ip, port, count *string) (*[]byte, *[]byte, *int, error) {
	if !((*ethInterface == "eth0") || (*ethInterface == "wlan0")) {
		return nil, nil, nil, errors.New("invalid interface")
	}
	ipv4NetIP := net.ParseIP(*ip)
	if ipv4NetIP == nil {
		return nil, nil, nil, errors.New("invalid ipv4 ip")
	}
	validIP4IP := strings.Contains(*ip, ":")
	if validIP4IP == true {
		return nil, nil, nil, errors.New("invalid ipv4 ip")
	}
	ipv4Byte := []byte(ipv4NetIP)
	portInt, _ := strconv.Atoi(*port) // convert first string octet to int
	if portInt < 1 || portInt > 65535 {
		return nil, nil, nil, errors.New("invalid port range")
	}
	portByte, _ := IntToByte(portInt)
	countInt, _ := strconv.Atoi(*count)
	if countInt < 1 || countInt > 1000000 {
		return nil, nil, nil, errors.New("invalid count range")
	}
	return &ipv4Byte, &portByte, &countInt, nil
}

// IPV6 converts a string into a properly formatted byte and int.
// It will return a []byte, []byte, int or an error if one occurred.
func IPV6(ethInterface, ip, port, count *string) (*[]byte, *[]byte, *int, error) {
	if !((*ethInterface == "eth0") || (*ethInterface == "wlan0")) {
		return nil, nil, nil, errors.New("invalid interface")
	}
	ipv6NetIP := net.ParseIP(*ip)
	if ipv6NetIP == nil {
		return nil, nil, nil, errors.New("invalid ip6 ip")
	}
	validIP6IP := strings.Contains(*ip, ".")
	if validIP6IP == true {
		return nil, nil, nil, errors.New("invalid ip6 ip")
	}
	ipv6Byte := []byte(ipv6NetIP)
	portInt, _ := strconv.Atoi(*port) // convert first string octet to int
	if portInt < 1 || portInt > 65535 {
		return nil, nil, nil, errors.New("invalid port range")
	}
	portByte, _ := IntToByte(portInt)
	countInt, _ := strconv.Atoi(*count)
	if countInt < 1 || countInt > 1000000 {
		return nil, nil, nil, errors.New("invalid count range")
	}
	return &ipv6Byte, &portByte, &countInt, nil
}
