package toy

import (
	"github.com/zerjioang/ssscomp/lib/common"
	"github.com/zerjioang/ssscomp/lib/math"
	"math/big"
)

const (
	// the bigger the prime size, the hardest factorization will be
	primeSize = 64
)

// let our homomorphic function
// E(m) = m + p * q
// m = plaintext value
// E(m) = encrypted value
// p = random number (big prime better, avoid factorization)
// q = random number (big prime better, avoid factorization)
type BigToyHomoScheme struct {
	common.HomomorphicSchema
	p *big.Int
	q *big.Int
}

func NewBigToyHomoScheme() *BigToyHomoScheme {
	scheme := new(BigToyHomoScheme)
	return scheme
}

func (s *BigToyHomoScheme) Generate() error {
	var pErr, qErr error
	s.p, pErr = math.Prime(primeSize)
	if pErr != nil {
		return pErr
	}
	s.q, qErr = math.Prime(primeSize)
	if qErr != nil {
		return qErr
	}
	return nil
}

// EncryptPadded as E(m) = m + p *q
func (s *BigToyHomoScheme) Encrypt(n int) *big.Int {
	cx := big.NewInt(0)
	cx.Mul(s.p, s.q)
	cx.Add(cx, big.NewInt(int64(n)))
	return cx
}

// decrypt as D(m) = c mod p
func (s *BigToyHomoScheme) Decrypt(cipher *big.Int) *big.Int {
	c := big.NewInt(0)
	c.Mod(cipher, s.p)
	return c
}
