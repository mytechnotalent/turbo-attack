// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package bit

import (
	"testing"
)

// Verify Clear produces a valid bit clear.
func TestClearProducesValidBitClear(t *testing.T) {
	// Params
	var n uint8
	n = 0b00000011
	var pos uint8
	pos = 1
	// Expect
	var want uint8
	want = 0b00000001
	// Calls
	got := Clear(&n, pos)
	// Asserts
	if want != *got {
		t.Fatalf(`Clear(&n, pos) = %v, want %v`, *got, want)
	}
}

// Verify Set produces a valid bit Set.
func TestSetProducesValidBitSet(t *testing.T) {
	// Params
	var n uint8
	n = 0b00000001
	var pos uint8
	pos = 1
	// Expect
	var want uint8
	want = 0b00000011
	// Calls
	got := Set(&n, pos)
	// Asserts
	if want != *got {
		t.Fatalf(`Set(&n, pos) = %v, want %v`, *got, want)
	}
}
