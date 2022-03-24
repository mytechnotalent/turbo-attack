// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// sudo code . --no-sandbox --user-data-dir=/home/anonymous/Documents
// usage: sudo -E /usr/local/go/bin/go test

package main

import (
	"go-test/convert"
	"go-test/routine"
	"testing"
)

// Verify TCP4 produces a valid TCP4 packet.
func TestTCP4ProducesValidTCP4Packet(t *testing.T) {
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "443"
	count := "1"
	ipv4Byte, portByte, _ := convert.IPV4(&ip, &port, &count)
	packetErr, writeErr := routine.IPv4(&ethInterface, ipv4Byte, portByte)
	if packetErr != nil {
		t.Errorf("packetErr")
	}
	if writeErr != nil {
		t.Errorf("writeErr")
	}
}

// Verify TCP6 produces a valid TCP6 packet.
func TestTCP6ProducesValidTCP6Packet(t *testing.T) {
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "1"
	ipv6Byte, portByte, _ := convert.IPV6(&ip, &port, &count)
	packetErr, writeErr := routine.IPv6(&ethInterface, ipv6Byte, portByte)
	if packetErr != nil {
		t.Errorf("packetErr")
	}
	if writeErr != nil {
		t.Errorf("writeErr")
	}
}
