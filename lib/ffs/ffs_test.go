package ffs

import (
	"math"
	"math/big"
	"testing"
)

func TestFFS(t *testing.T) {
	t.Run("big-int-example", func(t *testing.T) {
		p := big.NewInt(101) // random number
		q := big.NewInt(23)  // random number
		n := big.NewInt(0).Mul(p, q)
		peggy := secretBigShares{
			a: big.NewInt(5), // coprime to N (privKey)
			b: big.NewInt(7), // coprime to N (privKey)
			c: big.NewInt(3), // coprime to N (privKey)
		}
		t.Log("Peggy values: ", peggy)
		victor := secretBigShares{
			a: big.NewInt(1), // random
			b: big.NewInt(0), // random
			c: big.NewInt(1), // random
		}
		t.Log("Victor values: ", victor)
		// Next Peggy generates a random number (such as r=13).
		r := big.NewInt(13) // random < n
		peggyX := big.NewInt(0).Mul(r, r)
		peggyX = peggyX.Mod(peggyX, n)
		t.Log("Peggy x: ", peggyX)
	})
	t.Run("integer-example", func(t *testing.T) {
		p := 101 // random number
		q := 23  // random number
		n := p * q
		peggy := secretShares{
			a: 5, // coprime to N (privKey)
			b: 7, // coprime to N (privKey)
			c: 3, // coprime to N (privKey)
		}
		victor := secretShares{
			a: 1, // random
			b: 0, // random
			c: 1, // random
		}
		// Next Peggy generates a random number (such as r=13).
		r := 13 // random < n

		//She then calculates a value of x which is:
		// x = r² mod N
		peggyX := r * r % n
		// now Peggy send its X to Victor
		t.Log("Peggy x: ", peggyX)
		// Peggy makes the calculation
		peggyY := (r * pow(peggy.a, victor.a) * pow(peggy.b, victor.b) * pow(peggy.c, victor.c)) % n
		t.Log("y = ", peggyY)
		t.Log("y mod n = ", pow(peggyY, 2)%n)
		peggyV := secretShares{
			a: pow(peggy.a, 2) % n,
			b: pow(peggy.b, 2) % n,
			c: pow(peggy.c, 2) % n,
		}
		t.Log("peggy v values:", peggyV)
		t.Log("peggy sends values to victor")
		t.Log("victor computes...")
		victorY := (peggyX * pow(peggyV.a, victor.a) * pow(peggyV.b, victor.b) * pow(peggyV.c, victor.c)) % n
		y1prime := pow(peggyY, 2) % n
		t.Log("peggy y = ", peggyY)
		t.Log("victor y = ", victorY)
		valid := victorY == y1prime
		t.Log("valid proof: ", valid)
	})
}

func pow(base int, exponent int) int {
	r := math.Pow(float64(base), float64(exponent))
	return int(r)
}
