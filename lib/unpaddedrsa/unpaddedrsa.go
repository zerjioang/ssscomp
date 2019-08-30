package unpaddedrsa

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"math/big"
)

var (
	errPublicModulus       = errors.New("crypto/rsa: missing public modulus")
	errPublicExponentSmall = errors.New("crypto/rsa: public exponent too small")
	errPublicExponentLarge = errors.New("crypto/rsa: public exponent too large")
)

// checkPub sanity checks the public key before we use it.
// We require pub.E to fit into a 32-bit integer so that we
// do not have different behavior depending on whether
// int is 32 or 64 bits. See also
// https://www.imperialviolet.org/2012/03/16/rsae.html.
func checkPub(pub *rsa.PublicKey) error {
	if pub.N == nil {
		return errPublicModulus
	}
	if pub.E < 2 {
		return errPublicExponentSmall
	}
	if pub.E > 1<<31-1 {
		return errPublicExponentLarge
	}
	return nil
}

// Given a message x it is encrypted with the public keys
// it to get the ciphertext C(x) with:
// C(x)=x^e mod m
func EncryptUnpaddedRSA(pub *rsa.PublicKey, plaintext *big.Int) (*big.Int, error) {
	if err := checkPub(pub); err != nil {
		return nil, err
	}
	// encrypts using unpadded RSA
	// params
	// 1. plaintext
	// 2. exponent
	// 3. modulus
	// cx = encrypted result
	cx := new(big.Int)
	e := big.NewInt(int64(pub.E))
	cx.Exp(plaintext, e, pub.N)
	return cx, nil
}

// To decrypt a ciphertext C(x) one applies the private key:
//  m= C(x)^d mod m
func DecryptUnpaddedRSA(priv *rsa.PrivateKey, ciphertext *big.Int) (*big.Int, error) {
	if err := checkPub(&priv.PublicKey); err != nil {
		return nil, err
	}
	// decryption using unpadded RSA
	// params
	// 1. encrypted value
	// 2. private key
	// 3. modulus
	// c = decrypted result
	c := new(big.Int)
	c.Exp(ciphertext, priv.D, priv.N)
	return c, nil
}

func BigIntAsHex(v *big.Int) string {
	return fmt.Sprintf("%x", v) // or %X or upper case
}
func BigIntAsDecimal(v *big.Int) string {
	return fmt.Sprintf("%d", v) // or %X or upper case
}
