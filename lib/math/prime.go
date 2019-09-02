package math

import (
	"crypto/rand"
	"io"
	"math/big"
)

var (
	oneBig = big.NewInt(1)
	twoBig = big.NewInt(2)
)

// generates a prime number of size n
func Prime(size int) (*big.Int, error) {
	p, err := GeneratePrime(rand.Reader, 1, size)
	return p[0], err
}

// extracted from go/crypto/rsa/rsa.go
// modified to return a list of big prime numbers
func GeneratePrime(random io.Reader, nprimes int, bits int) ([]*big.Int, error) {
	primes := make([]*big.Int, nprimes)

NextSetOfPrimes:
	for {
		todo := bits
		// crypto/rand should set the top two bits in each prime.
		// Thus each prime has the form
		//   p_i = 2^bitlen(p_i) × 0.11... (in base 2).
		// And the product is:
		//   P = 2^todo × α
		// where α is the product of nprimes numbers of the form 0.11...
		//
		// If α < 1/2 (which can happen for nprimes > 2), we need to
		// shift todo to compensate for lost bits: the mean value of 0.11...
		// is 7/8, so todo + shift - nprimes * log2(7/8) ~= bits - 1/2
		// will give good results.
		if nprimes >= 7 {
			todo += (nprimes - 2) / 5
		}
		for i := 0; i < nprimes; i++ {
			var err error
			primes[i], err = rand.Prime(random, todo/(nprimes-i))
			if err != nil {
				return nil, err
			}
			todo -= primes[i].BitLen()
		}

		// Make sure that primes is pairwise unequal.
		for i, prime := range primes {
			for j := 0; j < i; j++ {
				if prime.Cmp(primes[j]) == 0 {
					continue NextSetOfPrimes
				}
			}
		}

		n := new(big.Int).Set(oneBig)
		totient := new(big.Int).Set(oneBig)
		pminus1 := new(big.Int)
		for _, prime := range primes {
			n.Mul(n, prime)
			pminus1.Sub(prime, oneBig)
			totient.Mul(totient, pminus1)
		}
		if n.BitLen() != bits {
			// This should never happen for nprimes == 2 because
			// crypto/rand should set the top two bits in each prime.
			// For nprimes > 2 we hope it does not happen often.
			continue NextSetOfPrimes
		}
		break
	}

	return primes, nil
}

func GenerateOddPrime(n *big.Int) *big.Int {
	// generate a 2n + 1. odd number
	p := big.NewInt(0)
	p.Mul(n, twoBig)
	p.Add(p, oneBig)
	return p
}
