package fhering

import (
	"fmt"
	"github.com/zerjioang/ssscomp/lib/common"
	"github.com/zerjioang/ssscomp/lib/math"
	"math/big"
)

// Fully homomorphic encryption in ring of binary integers

// fully homomorphic encryption: in ring Z2
type RingZtwo struct {

	// private key. aka secret parameter
	// p = 2k + 1
	p *big.Int

	// public keys
	publicKey []*big.Int
}

func NewRingZtwo() *RingZtwo {
	r := new(RingZtwo)
	return r
}

func (schema *RingZtwo) Generate() error {
	// generate 16 random bytes
	raw, err := common.GenerateRandomBytes(16)
	if err != nil {
		return err
	}
	// generate our random big number
	randomBig := big.NewInt(0)
	randomBig.SetBytes(raw)

	// generate a 2n + 1. odd number
	p := math.GenerateBigOdd(randomBig)

	fmt.Println(common.BigIntAsDecimal(p))
	schema.p = p
	// now its time to find z in Z2 ring by rule
	// z = 2r + m
	return nil
}

func (schema *RingZtwo) Encrypt() (*big.Int, error) {
	e := big.NewInt(0)
	m := big.NewInt(0)
	m.Rem(e, schema.p)
	return m, nil
}

func (schema *RingZtwo) Decrypt(ciphertext *big.Int) (*big.Int, error) {
	m := big.NewInt(0)
	m.Rem(ciphertext, schema.p)
	return m, nil
}
