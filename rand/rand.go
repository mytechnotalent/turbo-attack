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
	"errors"
	"math/big"
)

// Int generates a random int.
// It will return a valid *uint8 or error.
func Int(max int64) (*uint8, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return nil, errors.New("invalid uint8")
	}
	nInt64 := nBig.Int64()
	nUInt8 := uint8(nInt64)
	return &nUInt8, err
}

// Byte generates a random byte.
// It will return a valid byte or error.
func Byte(n int) ([]byte, error) {
	b := make([]byte, n, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
