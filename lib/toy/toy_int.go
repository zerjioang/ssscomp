package toy

import (
	"github.com/zerjioang/ssscomp/lib/common"
	"math"
)

// let our homomorphic function
// E(m) = m + p * q
// m = plaintext value
// E(m) = encrypted value
// p = random number (big prime better, avoid factorization)
// q = random number (big prime better, avoid factorization)
type IntegerToyHomoScheme struct {
	common.HomomorphicSchema
	p, q int
}

func NewIntegerToyHomoScheme() *IntegerToyHomoScheme {
	scheme := new(IntegerToyHomoScheme)
	return scheme
}

func (s *IntegerToyHomoScheme) Generate() error {
	s.p = int(math.Sqrt(float64(common.RandomInt() / 100000000000000)))
	s.q = int(math.Sqrt(float64(common.RandomInt() / 100000000000000)))
	return nil
}

// Encrypt as E(m) = m + p *q
func (s *IntegerToyHomoScheme) Encrypt(n int) int {
	return n + s.p*s.q
}

// Encrypt as E(m) = m + p *q
func (s *IntegerToyHomoScheme) EncryptF64(n int) float64 {
	return float64(s.Encrypt(n))
}

// decrypt as D(m) = c mod p
func (s *IntegerToyHomoScheme) Decrypt(c int) int {
	return c % s.p
}
