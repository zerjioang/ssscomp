package simple

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/s3go/lib/common"
	"testing"
)

func TestAdditiveShare(t *testing.T) {
	t.Run("3-participant-split", func(t *testing.T) {
		n1, err := SimpleAdditiveSecret(123456789, 3)
		assert.NotNil(t, n1)
		assert.Nil(t, err)
		fmt.Println(n1)

		//get shares for different participants
		alice, err := n1.Next()
		assert.NotNil(t, alice)
		assert.Nil(t, err)
		fmt.Println(alice)

		bob, err := n1.Next()
		assert.NotNil(t, bob)
		assert.Nil(t, err)
		fmt.Println(bob)

		mike, err := n1.Next()
		assert.NotNil(t, mike)
		assert.Nil(t, err)
		fmt.Println(mike)

		// since there are no more participants, next share request should fail
		charlie, err := n1.Next()
		assert.NotNil(t, charlie)
		assert.NotNil(t, err)

		assert.True(t, n1.Reconstruct([]common.Share{alice, bob, mike}))
		assert.False(t, n1.Reconstruct([]common.Share{alice, bob, mike}))
	})
	t.Run("3-participant-homomorphic-negation", func(t *testing.T) {
		// this test will compute a very basic homomorphic addition of two encrypted (unknown) values
		a, _ := SimpleAdditiveSecret(50, 3)
		t.Log(a)
		b := AdditiveNegation(a)
		assert.NotNil(t, b)
		t.Log(b)
		// once addition is done, execute the reconstruction
		assert.True(t, b.Reconstruct(-50))
	})
	t.Run("3-participant-homomorphic-addition", func(t *testing.T) {
		// this test will compute a very basic homomorphic addition of two encrypted (unknown) values
		a, _ := SimpleAdditiveSecret(50, 3)
		t.Log(a)
		b, _ := SimpleAdditiveSecret(30, 3)
		t.Log(b)
		c, err := AdditiveAdd(a, b)
		assert.Nil(t, err)
		assert.NotNil(t, c)
		t.Log(c)
		// once addition is done, execute the reconstruction
		assert.True(t, c.Reconstruct(80))
	})
	t.Run("3-participant-homomorphic-substraction", func(t *testing.T) {
		// this test will compute a very basic homomorphic addition of two encrypted (unknown) values
		a, _ := SimpleAdditiveSecret(50, 3)
		t.Log(a)
		b, _ := SimpleAdditiveSecret(30, 3)
		t.Log(b)
		c, err := AdditiveSubstraction(a, b)
		assert.Nil(t, err)
		assert.NotNil(t, c)
		t.Log(c)
		// once addition is done, execute the reconstruction
		assert.True(t, c.Reconstruct(20))
	})
	t.Run("3-participant-homomorphic-division", func(t *testing.T) {
		// this test will compute a very basic homomorphic addition of two encrypted (unknown) values
		a, _ := SimpleAdditiveSecret(60, 3)
		t.Log(a)
		b, _ := SimpleAdditiveSecret(30, 3)
		t.Log(b)
		c, err := AdditiveDivision(a, b)
		assert.Nil(t, err)
		assert.NotNil(t, c)
		t.Log(c)
		// once addition is done, execute the reconstruction
		assert.True(t, c.Reconstruct(2))
	})
	t.Run("3-participant-marshal", func(t *testing.T) {
		// this test will compute a very basic homomorphic addition of two encrypted (unknown) values
		a, _ := SimpleAdditiveSecret(50, 3)
		t.Log(a)
		b, _ := SimpleAdditiveSecret(30, 3)
		t.Log(b)
		c, err := AdditiveAdd(a, b)
		assert.Nil(t, err)
		assert.NotNil(t, c)
		t.Log(c)
		raw, _ := a.Json(0)
		t.Log(string(raw))
	})
	t.Run("2-homomorphic-mean", func(t *testing.T) {
		n1, _ := SimpleAdditiveSecret(50000, 3)
		n2, _ := SimpleAdditiveSecret(20000, 3)

		t.Log(n1)
		t.Log(n2)

		//get shares for different participants
		alice1, _ := n1.Next()
		bob1, _ := n1.Next()
		mike1, _ := n1.Next()

		alice2, _ := n2.Next()
		bob2, _ := n2.Next()
		mike2, _ := n2.Next()

		// each of the participants calculates the mean value
		t.Log("alice mean: ", alice2)
		t.Log("bob mean: ", bob2)
		t.Log("mike mean: ", mike2)

		t.Log(n1.Reconstruct([]common.Share{alice1, bob1, mike1}))
	})
	t.Run("2-participant-example", func(t *testing.T) {
		// this test will compute a very basic homomorphic addition of two encrypted (unknown) values
		// imagine we want to share our house value
		houseValue, _ := SimpleAdditiveSecret(200000, 2)
		//now lets share secret price value to the seller (me) and the possible buyer
		seller, _ := houseValue.Next()
		buyer, _ := houseValue.Next()
		fmt.Println("House value for seller:", seller)
		fmt.Println("House value for buyer:", buyer)

		// now lets create a secret value for what the buyer is willing to pay
		buyerPrice, _ := SimpleAdditiveSecret(150000, 2)
		seller2, _ := buyerPrice.Next()
		buyer2, _ := buyerPrice.Next()
		fmt.Println("Buyer value for seller:", seller2)
		fmt.Println("Buyer value for buyer:", buyer2)
		result, err := AdditiveComparison(houseValue, buyerPrice)
		assert.Nil(t, err)
		assert.NotNil(t, result)

		t.Log(result.Reconstruct(50000))
	})
}
