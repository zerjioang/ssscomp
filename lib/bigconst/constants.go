package bigconst

import "math/big"

var (
	BigE65537 = new(big.Int).SetInt64(65537)
	BigZero   = big.NewInt(0)
	BigOne    = big.NewInt(1)
)
