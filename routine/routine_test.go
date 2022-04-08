// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package routine

import (
	"testing"
)

// Verify IP4 produces a valid random TCP4 packet with random flags.
func TestIP4ProducesValidRandomTCP4PacketRandomFlags(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip4Byte := make([]byte, 16, 16)
	ip4Byte[0] = 0
	ip4Byte[1] = 0
	ip4Byte[2] = 0
	ip4Byte[3] = 0
	ip4Byte[4] = 0
	ip4Byte[5] = 0
	ip4Byte[6] = 0
	ip4Byte[7] = 0
	ip4Byte[8] = 0
	ip4Byte[9] = 0
	ip4Byte[10] = 0
	ip4Byte[11] = 0
	ip4Byte[12] = 192
	ip4Byte[13] = 168
	ip4Byte[14] = 0
	ip4Byte[15] = 2
	portByte := make([]byte, 2, 2)
	portByte[0] = 1
	portByte[1] = 187
	// Calls
	err := IP4(&ethInterface, &ip4Byte, &portByte)
	// Asserts
	if err != nil {
		t.Error(err)
	}
}

// Verify IP4 produces a invalid random TCP4 packet with random flags with invalid ethernet interface .
func TestIP4ProducesInvalidRandomTCP4PacketRandomFlagsInvalidEthInterface(t *testing.T) {
	// Params
	ethInterface := "foo"
	ip4Byte := make([]byte, 16, 16)
	ip4Byte[0] = 0
	ip4Byte[1] = 0
	ip4Byte[2] = 0
	ip4Byte[3] = 0
	ip4Byte[4] = 0
	ip4Byte[5] = 0
	ip4Byte[6] = 0
	ip4Byte[7] = 0
	ip4Byte[8] = 0
	ip4Byte[9] = 0
	ip4Byte[10] = 0
	ip4Byte[11] = 0
	ip4Byte[12] = 192
	ip4Byte[13] = 168
	ip4Byte[14] = 0
	ip4Byte[15] = 2
	portByte := make([]byte, 2, 2)
	portByte[0] = 1
	portByte[1] = 187
	// Calls
	err := IP4(&ethInterface, &ip4Byte, &portByte)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify IP6 produces a valid random TCP6 packet with random flags.
func TestIP6ProducesValidRandomTCP6PacketRandomFlags(t *testing.T) {
	// Params
	ethInterface := "eth0"
	ip6Byte := make([]byte, 16, 16)
	ip6Byte[0] = 254
	ip6Byte[1] = 128
	ip6Byte[2] = 0
	ip6Byte[3] = 0
	ip6Byte[4] = 0
	ip6Byte[5] = 0
	ip6Byte[6] = 0
	ip6Byte[7] = 0
	ip6Byte[8] = 0
	ip6Byte[9] = 0
	ip6Byte[10] = 0
	ip6Byte[11] = 0
	ip6Byte[12] = 0
	ip6Byte[13] = 0
	ip6Byte[14] = 0
	ip6Byte[15] = 0
	portByte := make([]byte, 2, 2)
	portByte[0] = 1
	portByte[1] = 187
	// Calls
	err := IP6(&ethInterface, &ip6Byte, &portByte)
	// Asserts
	if err != nil {
		t.Error(err)
	}
}

// Verify IP6 produces a invalid random TCP6 packet with random flags with invalid ethernet interface .
func TestIP6ProducesInvalidRandomTCP6PacketRandomFlagsInvalidEthInterface(t *testing.T) {
	// Params
	ethInterface := "foo"
	ip6Byte := make([]byte, 16, 16)
	ip6Byte[0] = 254
	ip6Byte[1] = 128
	ip6Byte[2] = 0
	ip6Byte[3] = 0
	ip6Byte[4] = 0
	ip6Byte[5] = 0
	ip6Byte[6] = 0
	ip6Byte[7] = 0
	ip6Byte[8] = 0
	ip6Byte[9] = 0
	ip6Byte[10] = 0
	ip6Byte[11] = 0
	ip6Byte[12] = 0
	ip6Byte[13] = 0
	ip6Byte[14] = 0
	ip6Byte[15] = 0
	portByte := make([]byte, 2, 2)
	portByte[0] = 1
	portByte[1] = 187
	// Calls
	err := IP6(&ethInterface, &ip6Byte, &portByte)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}
