// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides bit manipulation functionality.
package bit

import "errors"

// Clear clears a single bit in a uint8.
// It will return a valid *uint8 or an error if one occurred.
func Clear(n *uint8, pos uint8) (*uint8, error) {
	if *n < 0 || *n > 255 {
		return nil, errors.New("invalid uint8")
	}
	*n &= ^(1 << pos)
	return n, nil
}

// Set sets a single bit in a uint8.
// It will return a valid *uint8 or an error if one occurred.
func Set(n *uint8, pos uint8) (*uint8, error) {
	if *n < 0 || *n > 255 {
		return nil, errors.New("invalid uint8")
	}
	*n |= (1 << pos)
	return n, nil
}
