package bgn

import (
	"github.com/zerjioang/ssscomp/lib/common"
	"math/big"
)

func Generate() error {
	// choose large primes q,r and set n=qr
	q, err := common.Prime(5)
	if err != nil {
		return err
	}
	r, err := common.Prime(5)
	if err != nil {
		return err
	}
	// n = q * r
	n := big.NewInt(0)
	n.Mul(q, r)
}
