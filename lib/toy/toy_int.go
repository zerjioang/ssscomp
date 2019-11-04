package toy

import (
	"github.com/zerjioang/ssscomp/lib/common"
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
	s.p = common.GenerateSmallPrimeInt()
	s.q = common.GenerateSmallPrimeInt()
	return nil
}

func (s *IntegerToyHomoScheme) N() int {
	return s.p * s.q
}

// EncryptPadded as E(m) = m + p *q
func (s *IntegerToyHomoScheme) Encrypt(n int) int {
	return n + s.p*s.q
}

// EncryptPadded as E(m) = m + p *q
func (s *IntegerToyHomoScheme) EncryptF64(n int) float64 {
	return float64(s.Encrypt(n))
}

func (s *IntegerToyHomoScheme) EncryptF64F(n float64) float64 {
	return float64(s.Encrypt(int(n)))
}
func (s *IntegerToyHomoScheme) EncryptF64FArray(n []float64) []float64 {
	for i := range n {
		n[i] = s.EncryptF64F(n[i])
	}
	return n
}

// decrypt as D(m) = c mod p
func (s *IntegerToyHomoScheme) Decrypt(c int) int {
	return c % s.p
}
