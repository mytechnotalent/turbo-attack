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

// Int8ToByte converts an uint8 to a byte.
// It will return a *byte or an error if one occurred.
func Int8ToByte(n *uint8) (*byte, error) {
	if *n < 0 || *n > 255 {
		return nil, errors.New("invalid uint8")
	}
	nUint16 := uint16(*n)
	nByte := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(nByte, nUint16)
	nByte1 := nByte[1]
	return &nByte1, nil
}

// Int8ToByte converts an uint16 to a []byte.
// It will return a *[]byte or an error if one occurred.
func Int16ToByte(n *uint16) (*[]byte, error) {
	if *n < 0 || *n > 65535 {
		return nil, errors.New("invalid uint16")
	}
	nUint16 := uint16(*n)
	nByte := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(nByte, nUint16)
	return &nByte, nil
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
	portInt, err := strconv.Atoi(*port) // convert first string octet to int
	portUint16 := uint16(portInt)
	if err != nil {
		return nil, nil, nil, errors.New("invalid uint16")
	}
	if portInt < 1 || portInt > 65535 {
		return nil, nil, nil, errors.New("invalid port range")
	}
	portByte, _ := Int16ToByte(&portUint16)
	countInt, _ := strconv.Atoi(*count)
	if countInt < 1 || countInt > 2147483647 {
		return nil, nil, nil, errors.New("invalid count range")
	}
	return &ipv4Byte, portByte, &countInt, nil
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
	portUint16 := uint16(portInt)
	if portInt < 1 || portInt > 65535 {
		return nil, nil, nil, errors.New("invalid port range")
	}
	portByte, _ := Int16ToByte(&portUint16)
	countInt, _ := strconv.Atoi(*count)
	if countInt < 1 || countInt > 1000000 {
		return nil, nil, nil, errors.New("invalid count range")
	}
	return &ipv6Byte, portByte, &countInt, nil
}
