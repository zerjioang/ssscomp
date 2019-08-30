package simple

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/ssscomp/lib/common"
	"testing"
)

func TestAdditiveShare(t *testing.T) {
	t.Run("3-participant-share-recover", func(t *testing.T) {
		scheme, err := NewSimpleAdditiveScheme(3)
		assert.NotNil(t, scheme)
		assert.Nil(t, err)
		fmt.Println(scheme)
		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 20000
		shares := scheme.Generate(20000)

		//get shares for different participants
		alice := shares[0]
		bob := shares[1]
		mike := shares[2]

		r1, err := scheme.Reconstruct([]common.Shareable{alice, bob, mike})
		assert.NoError(t, err)
		assert.Equal(t, r1.IntValue(), 20000)
	})
	t.Run("3-participant-homomorphic-addition-ciphertext", func(t *testing.T) {
		x, _ := NewSimpleAdditiveScheme(3)
		fmt.Println(x)

		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 500
		sharesX := x.Generate(500)

		//get shares for different participants
		x1 := sharesX[0]
		x2 := sharesX[1]
		x3 := sharesX[2]

		t.Log(x1, x2, x3)

		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 100
		sharesY := x.Generate(100)
		y1 := sharesY[0]
		y2 := sharesY[1]
		y3 := sharesY[2]

		t.Log(y1, y2, y3)

		// created shares have homomorphic properties:
		// so that we can execute the addition as of x+y
		hsum1, _ := x1.Add(y1)
		hsum2, _ := x2.Add(y2)
		hsum3, _ := x3.Add(y3)

		t.Log(hsum1, hsum2, hsum3)
		result, _ := x.Reconstruct([]common.Shareable{hsum1, hsum2, hsum3})
		assert.Equal(t, result.IntValue(), 600)
	})
	t.Run("3-participant-homomorphic-linear-equation-ciphertext", func(t *testing.T) {
		x, _ := NewSimpleAdditiveScheme(3)
		fmt.Println(x)

		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 500
		sharesX := x.Generate(40)

		//get shares for different participants
		x1 := sharesX[0]
		x2 := sharesX[1]
		x3 := sharesX[2]

		t.Log(x1, x2, x3)

		//homomorphic linear equation ax + b
		a := 13
		hsum1, _ := x1.Mul(a)
		hsum2, _ := x2.Mul(a)
		hsum3, _ := x3.Mul(a)
		t.Log(hsum1, hsum2, hsum3)
		result, _ := x.Reconstruct([]common.Shareable{hsum1, hsum2, hsum3})
		assert.Equal(t, result.IntValue(), 520)

		//homomorphic addition
		b := common.NewIntSharePtr(2)
		hsum1, _ = hsum1.Add(b)
		hsum2, _ = hsum2.Add(b)
		hsum3, _ = hsum3.Add(b)

		t.Log(hsum1, hsum2, hsum3)
		result2, _ := x.Reconstruct([]common.Shareable{hsum1, hsum2, hsum3})
		assert.Equal(t, result2.IntValue(), 525)
	})
	t.Run("3-participant-homomorphic-multiply-plaintext", func(t *testing.T) {
		x, _ := NewSimpleAdditiveScheme(3)

		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 50
		sharesY := x.Generate(50)
		y1 := sharesY[0]
		y2 := sharesY[1]
		y3 := sharesY[2]

		t.Log(y1, y2, y3)

		//homomorphic multiplication of ciphertext and plaintext
		hsum1, _ := y1.Mul(3)
		hsum2, _ := y2.Mul(3)
		hsum3, _ := y3.Mul(3)

		t.Log(hsum1, hsum2, hsum3)
		result, _ := x.Reconstruct([]common.Shareable{hsum1, hsum2, hsum3})
		assert.Equal(t, result.IntValue(), 150)
	})
	t.Run("3-participant-homomorphic-negation-plaintext", func(t *testing.T) {
		x, _ := NewSimpleAdditiveScheme(3)

		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 100
		sharesY := x.Generate(50)
		y1 := sharesY[0]
		y2 := sharesY[1]
		y3 := sharesY[2]

		t.Log(y1, y2, y3)

		//homomorphic negation of shares
		hsum1, _ := y1.Neg()
		hsum2, _ := y2.Neg()
		hsum3, _ := y3.Neg()

		t.Log(hsum1, hsum2, hsum3)
		result, _ := x.Reconstruct([]common.Shareable{hsum1, hsum2, hsum3})
		assert.Equal(t, result.IntValue(), -50)
	})
	t.Run("3-participant-homomorphic-division-plaintext", func(t *testing.T) {
		x, _ := NewSimpleAdditiveScheme(3)
		fmt.Println(x)

		//use additive secret sharing for new data distribution among 3 participants
		// in this example, data will be value 500
		sharesX := x.Generate(500)

		//get shares for different participants
		x1 := sharesX[0]
		x2 := sharesX[1]
		x3 := sharesX[2]

		t.Log(x1, x2, x3)

		//homomorphic division of plaintext. divide by 2
		n2 := common.NewIntSharePtr(2)
		hsum1, _ := x1.Div(n2)
		hsum2, _ := x2.Div(n2)
		hsum3, _ := x3.Div(n2)

		t.Log(hsum1, hsum2, hsum3)
		result, _ := x.Reconstruct([]common.Shareable{hsum1, hsum2, hsum3})
		assert.Equal(t, result.IntValue(), 250)
	})

}
