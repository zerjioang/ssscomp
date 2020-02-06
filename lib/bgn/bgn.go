package bgn

import (
	"github.com/zerjioang/ssscomp/lib/bigconst"
	"github.com/zerjioang/ssscomp/lib/math"
)

func Generate() error {
	// choose large primes q,r and set n=qr
	q, err := math.Prime(5)
	if err != nil {
		return err
	}
	r, err := math.Prime(5)
	if err != nil {
		return err
	}
	// n = q * r
	n := bigconst.BigZero
	n.Mul(q, r)
}
