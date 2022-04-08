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

// Verify TCP4 produces an invalid TCP4 packet by passing in an invalid IP.
func TestTCP4ProducesInvalidTCP4PacketWithInvalidIP(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip := "foo"
	port := "443"
	count := "1"
	// Calls
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	count := "1000001"
	// Calls
	_, _, _, err := IPV4(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
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
	count := "1000001"
	// Calls
	_, _, _, err := IPV6(&ethInterface, &ip, &port, &count)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}
