package ffs

import "math/big"

// The user has three secret numbers of a=5; b=7; and c=3.
//Note that these numbers need to be co-prime to N, so that
//they do not share a factor with N. If we choose N as the
//multiplication of two prime numbers, we only need to avoid these numbers.
type secretShares struct {
	a int
	b int
	c int
}

type secretBigShares struct {
	a *big.Int
	b *big.Int
	c *big.Int
}
