// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// sudo code . --no-sandbox --user-data-dir=/home/anonymous/Documents/user-data-dir

package convert

import (
	"testing"
)

// Verify IntToByte produces a valid conversion.
func TestIntToByteProducesValidConversion(t *testing.T) {
	// Params
	n := 42
	// Calls
	_, err := IntToByte(n)
	// Asserts
	if err != nil {
		t.Error(err)
	}
}

// Verify IntToByte produces an invalid IntToByte []byte by passing in an invalid lower int range.
func TestIntToByteProducesInvalidConversionWithInvalidLowerIntRange(t *testing.T) {
	// Params
	n := -1
	// Calls
	_, err := IntToByte(n)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify IntToByte produces an invalid IntToByte []byte by passing in an invalid higher int range.
func TestIntToByteProducesInvalidConversionWithInvalidHigherIntRange(t *testing.T) {
	// Params
	n := 1000001
	// Calls
	_, err := IntToByte(n)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces a valid TCP4 packet.
func TestTCP4ProducesValidTCP4Packet(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err != nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid ethInterface.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidEthInterface(t *testing.T) {
	// Params
	ethInterface := "foo"
	ip := "192.168.0.2"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}
