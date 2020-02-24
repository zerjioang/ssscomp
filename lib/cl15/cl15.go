package main

import (
	"errors"
	"math/big"

	"github.com/zerjioang/ssscomp/lib/bigconst"
)

func generateIds(p, q, uid, xa, xb *big.Int, bits int) error {
	if p == nil {
		var err error
		p, err = bigconst.LargePrime(bits)
		if err != nil {
			return err
		}
	}

	if q == nil {
		var err error
		q, err = bigconst.LargePrime(bits)
		if err != nil {
			return err
		}
	}
	// n = p * q
	n := new(big.Int).Mul(p, q)

	phiL := new(big.Int).Sub(p, bigconst.BigOne)
	phiR := new(big.Int).Sub(q, bigconst.BigOne)

	// PHI=(p-1)*(q-1)
	phi := new(big.Int).Mul(phiL, phiR)

	// d=inverse_of(e,PHI)
	d := new(big.Int).ModInverse(bigconst.BigE65537, phi)

	id1 := new(big.Int).Exp(uid, xa, n)
	id2 := new(big.Int).Exp(uid, xb, n)
	if id1 == nil {
		return errors.New("failed to create id1")
	}
	if id2 == nil {
		return errors.New("failed to create id2")
	}

	cipherOne := new(big.Int).Exp(id1, bigconst.BigE65537, n)

	val := new(big.Int).Exp(cipherOne, xb, n)
	modInvxaphi := new(big.Int).ModInverse(xa, phi)
	if modInvxaphi == nil {
		return errors.New("failed to compute modInvxaphi")
	}

	val = new(big.Int).Exp(val, modInvxaphi, n)
	val = new(big.Int).Exp(val, d, n)
	return nil
}
