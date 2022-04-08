// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package convert

import (
	"bytes"
	"testing"
)

// Verify Int8ToByte produces a valid conversion.
func TestInt8ToByteProducesValidConversion(t *testing.T) {
	// Params
	var n uint8
	n = 42
	// Expect
	want := byte(42)
	// Calls
	got := Int8ToByte(&n)
	// Asserts
	if want != *got {
		t.Fatalf(`Int8ToByte(&n) = %v, want %v`, *got, want)
	}
}

// Verify Int16ToByte produces a valid conversion.
func TestInt16ToByteProducesValidConversion(t *testing.T) {
	// Params
	var n uint16
	n = 42
	// Expect
	want := make([]byte, 2, 2)
	want[0] = 0
	want[1] = 42
	// Calls
	got := Int16ToByte(&n)
	// Asserts
	if !bytes.Equal(want, *got) {
		t.Fatalf(`Int16ToByte(&n) = %v, want %v`, *got, want)
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
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err != nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid IP.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidIP(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "foo"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in a valid IPV6 IP.
func TestTCP4ProducesInvalidTCP4PacketWithValidIPV6IP(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
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
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid port.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidPort(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "foo"
	count := "1"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid port lower int range.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidPortLowerIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "0"
	count := "1"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid port higher int range.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidPortHigherIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "65536"
	count := "1"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid count.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidCount(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "443"
	count := "foo"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid count lower int range.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidCountLowerIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "443"
	count := "0"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid count higher int range.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidCountHigherIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "443"
	count := "2147483648"
	// Calls
	_, _, _, err := IP4(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces a valid TCP6 packet.
func TestTCP6ProducesValidTCP6Packet(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err != nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid IP.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidIP(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "foo"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in a valid IPV4 IP.
func TestTCP6ProducesInvalidTCP6PacketWithValidIPV4IP(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "192.168.0.2"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid ethInterface.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidEthInterface(t *testing.T) {
	// Params
	ethInterface := "foo"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid port.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidPort(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "foo"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid port lower int range.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidPortLowerIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "0"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid port higher int range.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidPortHigherIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "65536"
	count := "1"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid count.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidCount(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "foo"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid count lower int range.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidCountLowerIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "0"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify TCP6 produces an invalid TCP6 packet by passing in an invalid count higher int range.
func TestTCP6ProducesInvalidTCP6PacketWithInvalidCountHigherIntRange(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "fe80:0000:0000:0000:0000:0000:0000:0002"
	port := "443"
	count := "2147483648"
	// Calls
	_, _, _, err := IP6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}
