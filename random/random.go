// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides random number generation.
package random

import (
	"crypto/rand"
	"math/big"
)

// Int generates a random int.
// It will return a valid int or error.
func Int(max int64) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	nInt64 := nBig.Int64()
	nInt := int(nInt64)
	return nInt, err
}

// Byte generates a random byte.
// It will return a valid byte or error.
func Byte(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
