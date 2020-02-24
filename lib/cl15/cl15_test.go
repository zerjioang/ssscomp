package main

import (
	"log"
	"math/big"
	"testing"

	"github.com/zerjioang/ssscomp/lib/bigconst"
)

const (
	//default number of bits used
	bits = 2048
)

func ExampleCL15() {
	p, err := bigconst.LargePrime(bits)
	if err != nil {
		log.Fatal("could not generate prime p")
	}

	q, err := bigconst.LargePrime(bits)
	if err != nil {
		log.Fatal("could not generate prime q")
	}

	n := new(big.Int).Mul(p, q)

	e := new(big.Int).SetInt64(65537)

	phiL := new(big.Int).Sub(p, big.NewInt(1))
	phiR := new(big.Int).Sub(q, big.NewInt(1))

	phi := new(big.Int).Mul(phiL, phiR)

	d := new(big.Int).ModInverse(e, phi)

	uid := new(big.Int).SetInt64(19382983298)
	xa := new(big.Int).SetInt64(3929333233)
	xb := new(big.Int).SetInt64(2389239238)

	id1 := new(big.Int).Exp(uid, xa, n)
	id2 := new(big.Int).Exp(uid, xb, n)

	log.Println("id 1", id1)
	log.Println("id 2", id2)
	log.Println("xa", xa)
	log.Println("xb", xb)
	log.Println("e", e)
	log.Println("N", n)
	log.Println("p", p)
	log.Println("q", q)
	log.Println("RSA Encryption parameters. Pubkey [e, N]")

	cipherOne := new(big.Int).Exp(id1, e, n)
	log.Println("Derived ID1:", id1)
	val := new(big.Int).Exp(cipherOne, xb, n)
	modInvxaphi := new(big.Int).ModInverse(xa, phi)
	log.Println(val)
	if modInvxaphi == nil {
		log.Fatalf("failed to compute modInvxaphi")
	}
	log.Println(modInvxaphi)
	log.Println(n)

	val = new(big.Int).Exp(val, modInvxaphi, n)
	val = new(big.Int).Exp(val, d, n)
	log.Println("Derived ID2:", val)
}

func TestCL15(t *testing.T) {
	t.Run("example-1", func(t *testing.T) {
		ExampleCL15()
	})
	t.Run("example-2", func(t *testing.T) {
		_ = generateIds(
			nil,
			nil,
			new(big.Int).SetInt64(19382983298),
			new(big.Int).SetInt64(3929333233),
			new(big.Int).SetInt64(2389239238),
			bits,
		)
	})
	t.Run("profiling", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			_ = generateIds(
				nil,
				nil,
				new(big.Int).SetInt64(19382983298),
				new(big.Int).SetInt64(3929333233),
				new(big.Int).SetInt64(2389239238),
				bits,
			)
		}
	})
}
