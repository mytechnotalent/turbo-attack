// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package rand

import (
	"testing"
)

// Verify UInt8 produces a valid a random uint8.
func TestUInt8ProducesValidRandomUInt8(t *testing.T) {
	// Params
	var n int64
	n = 255
	// Expect
	var want uint8
	want = 255
	// Calls
	got := UInt8(n)
	// Asserts
	if *got > want {
		t.Fatalf(`UInt8(n) = %v, want %v which is greater than n`, *got, want)
	}
}

// Verify Byte produces a valid a random byte.
func TestByteProducesValidRandomByte(t *testing.T) {
	// Params
	var n uint8
	n = 255
	// Expect
	want := []byte{255}
	// Calls
	got, _ := Byte(n)
	// Asserts
	if got[0] > want[0] {
		t.Fatalf(`Int(n) = %v, want %v which is greater than n`, got, want)
	}
}
