package unpaddedrsa

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/ssscomp/lib/common"
	"math/big"
	"testing"
)

func TestUnpaddedRsa(t *testing.T) {
	t.Run("homomorphic-multiplication", func(t *testing.T) {

		// generate key
		priv, pub, _ := GenerateKeyPair(1024)
		assert.NotNil(t, priv)
		assert.NotNil(t, pub)

		// given a RSA public key
		// pub key(n, e)
		t.Log("given RSA public key: ", pub.N, pub.E, pub.Size())
		// and given plain text x1 and x2
		x1 := 200
		x2 := 3
		// we can compute
		// enc(x1) * enc(x2) == ( x1^e * x2^e ) mod n
		b1 := big.NewInt(int64(x1))
		b2 := big.NewInt(int64(x2))

		// unpadded RSA encrypted version of x1
		cipher1, err := EncryptUnpaddedRSA(pub, b1)
		assert.NoError(t, err)
		assert.NotNil(t, cipher1)
		t.Log("encrypted x1: ", common.BigIntAsHex(cipher1))
		t.Log("encrypted x1: ", common.BigIntAsDecimal(cipher1))

		// unpadded RSA encrypted version of x2
		cipher2, err := EncryptUnpaddedRSA(pub, b2)
		assert.NoError(t, err)
		assert.NotNil(t, cipher2)
		t.Log("encrypted x2: ", common.BigIntAsHex(cipher2))
		t.Log("encrypted x2: ", common.BigIntAsDecimal(cipher2))

		// compute homomorphic multiplication over encrypted values cipher1 and cipher2
		c12 := big.NewInt(int64(0))
		c12.Mul(cipher1, cipher2)
		// mod n?

		decrypted12, err := DecryptUnpaddedRSA(priv, c12)
		assert.NoError(t, err)
		assert.NotNil(t, decrypted12)
		t.Log("decrypted value: ", decrypted12.String())
		assert.Equal(t, decrypted12.String(), "600")
	})
}

func ExampleHomomorphicRsaMul() {
	// generate key
	priv, pub, _ := GenerateKeyPair(1024)
	// and given plain text x1 and x2
	x1 := 200
	x2 := 3
	// we can compute
	// enc(x1) * enc(x2) == ( x1^e * x2^e ) mod n
	b1 := big.NewInt(int64(x1))
	b2 := big.NewInt(int64(x2))

	// unpadded RSA encrypted version of x1
	cipher1, _ := EncryptUnpaddedRSA(pub, b1)
	fmt.Println("encrypted x1: ", common.BigIntAsHex(cipher1))

	// unpadded RSA encrypted version of x2
	cipher2, _ := EncryptUnpaddedRSA(pub, b2)
	fmt.Println("encrypted x2: ", common.BigIntAsHex(cipher2))

	// compute homomorphic multiplication over encrypted values cipher1 and cipher2
	c12 := big.NewInt(int64(0))
	c12.Mul(cipher1, cipher2)

	decrypted12, _ := DecryptUnpaddedRSA(priv, c12)
	plainMul := decrypted12.String()
	fmt.Println("decrypted value: ", plainMul)
}
