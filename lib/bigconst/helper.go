package bigconst

import (
	"crypto/rand"
	"math/big"
)

func LargePrime(k int) (p *big.Int, err error) {
	return rand.Prime(rand.Reader, k)
}
