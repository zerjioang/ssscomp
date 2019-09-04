package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func BigIntAsHex(v *big.Int) string {
	return fmt.Sprintf("%x", v) // or %X or upper case
}
func BigIntAsDecimal(v *big.Int) string {
	return fmt.Sprintf("%d", v) // or %X or upper case
}

// generates a random big interger from 0 to Q - 1
// returns a uniform random value in [0, max). It panics if max <= 0.
func RandomBigInteger(max *big.Int) (*big.Int, error) {
	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	return n, err
}