package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleHomoCipher(t *testing.T) {
	t.Run("example-0", func(t *testing.T) {
		c := NewSimpleHomoCipher(0)
		encrypted, err := c.Do()
		assert.NoError(t, err)
		assert.NotNil(t, encrypted)
		t.Log(c)
		t.Log(encrypted)
		plain := c.Recover(encrypted)
		t.Log(plain)
		assert.Equal(t, plain, 0)
	})
	t.Run("example-1", func(t *testing.T) {
		c := NewSimpleHomoCipher(1)
		encrypted, err := c.Do()
		assert.NoError(t, err)
		assert.NotNil(t, encrypted)
		t.Log(c)
		t.Log(encrypted)
		plain := c.Recover(encrypted)
		t.Log(plain)
		assert.Equal(t, plain, 1)
	})

	t.Run("homomorphic-addition", func(t *testing.T) {
		a := NewSimpleHomoCipher(1)
		encA, _ := a.Do()
		t.Log(a)

		b := NewSimpleHomoCipher(1)
		encB, _ := b.Do()
		t.Log(b)

		sum := encA + encB
		t.Log("homomorphic addition: ", sum)
		plainSum := b.Recover(sum)
		t.Log("homomorphic addition as plain text: ", plainSum)
	})
}
