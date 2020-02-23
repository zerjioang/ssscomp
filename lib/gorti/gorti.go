package gorti

import (
	"github.com/zerjioang/ssscomp/lib/common"
	"math/big"
)

// Define the encryption key k = (p, q,m,r)
// Q will be kept secret (It is secret key)
// That number m is also a secret key to encrypt the data
// it is generally regarded that m should be at least 1024, if not 2048
type EncryptionKey struct {
	p, q, m, r *big.Int // m is the public key
}

// Ek(X)=)(mod m)
// Encrypt(X,m,p,q,r)
// Assume X € Zp
// Compute
// Y = (X + r*p q) (mod m)
// Output Y € Zc
// PROOF
// Proof : Encrypt the message X
// E(X) = )(mod m)
// Cipher text Y will be (X+rp)
func (e *EncryptionKey) encrypt(x *big.Int) *big.Int {
	pr := big.NewInt(0).Mul(e.p, e.r)
	return big.NewInt(0).Add(x, pr)
}

// Decryption will be done with the secret key k = ( p ),
// X= Dk(Y)= C mod p. But can be broken if p can be
// discovered but which is a very tough to solve.
// Decrypt Y = X = Y mod p
// 		= (X+rp) mod p
// 		= rp mod p + X mod p
// 		= X Plaintext
func (e *EncryptionKey) decrypt(y *big.Int) *big.Int {
	return big.NewInt(0).Mod(y, e.p)
}

func NewGorti(size uint) *EncryptionKey {
	// A large prime number ‘p’
	p, _ := common.GeneratePrime(size)
	// another prime number ‘q’ such that q < p
	q, _ := common.RandomPrimeLess(p)
	// a random number ‘r’ has taken to make the scheme non deterministic
	r, _ := common.RandomBigLess(p)
	// calculate m as m = p*q for modulo m operations
	m := big.NewInt(0).Mul(p, q)
	return &EncryptionKey{
		p: p,
		q: q,
		r: r,
		m: m,
	}
}
