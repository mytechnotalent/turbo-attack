// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides bit manipulation functionality.
package bit

// Clear clears a single bit in an int.
// It will return an int with the respective bit cleared.
func Clear(n int, pos uint) int {
	n &= ^(1 << pos)
	return n
}

// Set sets a single bit in an int.
// It will return an int with the respective bit set.
func Set(n int, pos uint) int {
	n |= (1 << pos)
	return n
}
