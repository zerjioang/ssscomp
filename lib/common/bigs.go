package common

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
)

func BigIntAsHex(v *big.Int) string {
	return fmt.Sprintf("%x", v) // or %X or upper case
}
func BigIntAsDecimal(v *big.Int) string {
	return fmt.Sprintf("%d", v)
}
func FromSafeHex(hexBigInt string) *big.Int {
	ret, err := BigIntFromHex(hexBigInt)
	if err != nil {
		panic(err)
	}
	return ret
}

func BigIntFromHex(hex string) (*big.Int, error) {
	n, err := new(big.Int).SetString(hex, 16)
	if !err {
		msg := fmt.Sprintf("Cannot convert %s to int as hexadecimal", hex)
		return nil, errors.New(msg)
	}
	return n, nil
}

// generates a random big interger from 0 to Q - 1
// returns a uniform random value in [0, max). It panics if max <= 0.
func RandomBigInteger(max *big.Int) (*big.Int, error) {
	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	return n, err
}

//  returns n! = n*(n-1)*(n-2)...3*2*1
func Factorial(n int) *big.Int {
	ret := big.NewInt(1)
	for i := 1; i <= n; i++ {
		ret = new(big.Int).Mul(ret, big.NewInt(int64(i)))
	}
	return ret
}

//  Returns 2 primes such that p = 2 * q + 1 and that the length of
//  p is nbits.  `p` is called a safe prime
func GenerateSafePrimes(nbits int, random io.Reader) (p, q *big.Int, err error) {
	for {
		q, err = rand.Prime(random, nbits-1)
		if err != nil {
			return
		}
		p = (new(big.Int)).Mul(q, big.NewInt(2))
		p = p.Add(p, big.NewInt(1))
		if p.ProbablyPrime(50) { //a probability of 2**-100 of not being prime
			return
		}
	}
}

func GenerateSmallPrimeInt() int {
	p, _ := GeneratePrime(16)
	return int(p.Int64())
}

func GeneratePrime(size uint) (p *big.Int, err error) {
	return rand.Prime(rand.Reader, int(size))
}

// Generate a random element in the group of all the elements in Z/nZ that
// has a multiplicative inverse.
func GetRandomNumberInMultiplicativeGroup(n *big.Int, random io.Reader) (*big.Int, error) {
	r, err := rand.Int(random, n)
	if err != nil {
		return nil, err
	}
	zero := big.NewInt(0)
	one := big.NewInt(1)
	if zero.Cmp(r) == 0 || one.Cmp(new(big.Int).GCD(nil, nil, n, r)) != 0 {
		return GetRandomNumberInMultiplicativeGroup(n, random)
	}
	return r, nil

}

//  Return a random generator of RQn with high probability.  THIS METHOD
//  ONLY WORKS IF N IS THE PRODUCT OF TWO SAFE PRIMES! This heuristic is used
//  threshold signature paper in the Victor Shoup
func GetRandomGeneratorOfTheQuadraticResidue(n *big.Int, rand io.Reader) (*big.Int, error) {
	r, err := GetRandomNumberInMultiplicativeGroup(n, rand)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Mod(new(big.Int).Mul(r, r), n), nil
}
