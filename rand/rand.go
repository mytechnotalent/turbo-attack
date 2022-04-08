// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package rand provides random number generation.
package rand

import (
	"crypto/rand"
	"math/big"
)

// UInt8 generates a random uint8.
// It will return a *uint8.
func UInt8(max int64) *uint8 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(max))
	nInt64 := nBig.Int64()
	nUInt8 := uint8(nInt64)
	return &nUInt8
}

// Byte generates a random byte.
// It will return a valid byte or error.
func Byte(n uint8) ([]byte, error) {
	b := make([]byte, n, n)
	_, err := rand.Read(b)
	return b, err
}
