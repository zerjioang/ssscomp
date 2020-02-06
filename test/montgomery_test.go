package test

import (
	"fmt"
	"math/big"
	"testing"
)

type mont struct {
	n  uint     // m.BitLen()
	m  *big.Int // modulus, must be odd
	r2 *big.Int // (1<<2n) mod m
}

func newMont(m *big.Int) *mont {
	if m.Bit(0) != 1 {
		return nil
	}
	n := uint(m.BitLen())
	x := big.NewInt(1)
	x.Sub(x.Lsh(x, n), m)
	return &mont{n, new(big.Int).Set(m), x.Mod(x.Mul(x, x), m)}
}

// Montgomery reduction algorithm
func (m mont) reduce(t *big.Int) *big.Int {
	a := new(big.Int).Set(t)
	for i := uint(0); i < m.n; i++ {
		if a.Bit(0) == 1 {
			a.Add(a, m.m)
		}
		a.Rsh(a, 1)
	}
	if a.Cmp(m.m) >= 0 {
		a.Sub(a, m.m)
	}
	return a
}

func TestMontgometry(t *testing.T) {

	p := 130001
	a := 50000
	b := 60000

	m := big.NewInt(int64(p))
	fmt.Printf("a=%d, b=%d, p=%d\n\n", a, b, p)

	mr := newMont(m)

	x1 := big.NewInt(int64(a))
	x2 := big.NewInt(int64(b))
	t1 := mr.reduce(x1)
	t2 := mr.reduce(x2)
	res := mr.reduce(new(big.Int).Mul(t1, t2))

	fmt.Printf("Value: %s*%s (mod %s) = %s\n\n", x1, x2, m, res)
	prod := mr.reduce(mr.r2)
	base := mr.reduce(t1.Mul(x1, mr.r2))
	exp := new(big.Int).Set(x2)
	for exp.BitLen() > 0 {
		if exp.Bit(0) == 1 {
			prod = mr.reduce(prod.Mul(prod, base))
		}
		exp.Rsh(exp, 1)
		base = mr.reduce(base.Mul(base, base))
	}
	fmt.Printf("\nValue: %s^%s (mod %s) = %s\n", x1, x2, m, mr.reduce(prod))
	fmt.Printf("Traditional method result = %s", new(big.Int).Exp(x1, x2, m))
}
