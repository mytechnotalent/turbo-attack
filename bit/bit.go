// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides bit manipulation functionality.
package bit

// Clear clears a single bit in a uint8.
// It will return a *uint8.
func Clear(n *uint8, pos uint8) *uint8 {
	*n &= ^(1 << pos)
	return n
}

// Set sets a single bit in a uint8.
// It will return a *uint8.
func Set(n *uint8, pos uint8) *uint8 {
	*n |= (1 << pos)
	return n
}
