package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	//default number of bits used
	bits = 2048
)

var (
	bigE65537 = new(big.Int).SetInt64(65537)
	bigOne = big.NewInt(1)
)
func generateLargePrime(k int) (p *big.Int, err error) {
	return rand.Prime(rand.Reader, k)
}

func generateIds(p, q, uid, xa, xb *big.Int) error {
	if p == nil {
		var err error
		p, err = generateLargePrime(bits)
		if err != nil {
			return err
		}
	}

	if q == nil {
		var err error
		q, err = generateLargePrime(bits)
		if err != nil {
			return err
		}
	}
	// n=p*q
	n := new(big.Int).Mul(p, q)

	phiL := new(big.Int).Sub(p, bigOne)
	phiR := new(big.Int).Sub(q, bigOne)
	// PHI=(p-1)*(q-1)
	phi := new(big.Int).Mul(phiL, phiR)
	// d=inverse_of(e,PHI)
	d := new(big.Int).ModInverse(bigE65537, phi)

	id1 := new(big.Int).Exp(uid, xa, n)
	id2 := new(big.Int).Exp(uid, xb, n)
	if id1 == nil {
		return errors.New("failed to create id1")
	}
	if id2 == nil {
		return errors.New("failed to create id2")
	}

	cipherOne := new(big.Int).Exp(id1, bigE65537, n)

	val := new(big.Int).Exp(cipherOne, xb, n)
	modInvxaphi := new(big.Int).ModInverse(xa, phi)
	if modInvxaphi == nil {
		return errors.New("failed to compute modInvxaphi")
	}

	val = new(big.Int).Exp(val, modInvxaphi, n)
	val = new(big.Int).Exp(val, d, n)
	return nil
}