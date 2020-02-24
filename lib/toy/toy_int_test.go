package toy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerToyHomoScheme(t *testing.T) {
	t.Run("encrypt-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		encryptedValue := scheme.Encrypt(3)
		assert.NotNil(t, encryptedValue)
		assert.NotEqual(t, encryptedValue, 3)
		t.Log(encryptedValue)

		//decrypt
		plaintext := scheme.Decrypt(encryptedValue)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 3)
		t.Log(plaintext)
	})
	t.Run("encrypt-decrypt-example", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := IntegerToyHomoScheme{p: 20, q: 9}

		//encrypt
		encryptedValue := scheme.Encrypt(3)
		assert.NotNil(t, encryptedValue)
		assert.NotEqual(t, encryptedValue, 3)
		t.Log(encryptedValue)

		//decrypt
		plaintext := scheme.Decrypt(encryptedValue)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 3)
		t.Log(plaintext)
	})
	t.Run("encrypt-add-decrypt-example", func(t *testing.T) {
		scheme := IntegerToyHomoScheme{p: 20, q: 9}
		//encrypt
		a := scheme.Encrypt(3)
		b := scheme.Encrypt(6)
		c := a + b
		//decrypt
		t.Log(a, b, c)
		plaintext := scheme.Decrypt(c)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 9)
		t.Log(plaintext)
	})
	t.Run("encrypt-mul-decrypt-example", func(t *testing.T) {
		scheme := IntegerToyHomoScheme{p: 20, q: 9}
		//encrypt
		a := scheme.Encrypt(3)
		b := scheme.Encrypt(6)
		c := a * b
		t.Log(a, b, c)
		//decrypt
		plaintext := scheme.Decrypt(c)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 18)
		t.Log(plaintext)
	})
	t.Run("encrypt-add-plaintext-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		encryptedValue := scheme.Encrypt(3)
		assert.NotNil(t, encryptedValue)
		assert.NotEqual(t, encryptedValue, 3)
		t.Log(encryptedValue)

		encryptedValue += 5

		//decrypt
		plaintext := scheme.Decrypt(encryptedValue)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 8)
		t.Log(plaintext)
	})
	t.Run("encrypt-sub-plaintext-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		encryptedValue := scheme.Encrypt(3)
		assert.NotNil(t, encryptedValue)
		assert.NotEqual(t, encryptedValue, 3)
		t.Log(encryptedValue)

		encryptedValue -= 2

		//decrypt
		plaintext := scheme.Decrypt(encryptedValue)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 1)
		t.Log(plaintext)
	})
	t.Run("encrypt-mul-plaintext-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		encryptedValue := scheme.Encrypt(3)
		assert.NotNil(t, encryptedValue)
		assert.NotEqual(t, encryptedValue, 3)
		t.Log(encryptedValue)

		encryptedValue *= 2

		//decrypt
		plaintext := scheme.Decrypt(encryptedValue)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 6)
		t.Log(plaintext)
	})
	t.Run("encrypt-div-plaintext-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		encryptedValue := scheme.Encrypt(20)
		assert.NotNil(t, encryptedValue)
		assert.NotEqual(t, encryptedValue, 20)
		t.Log(encryptedValue)

		encryptedValue /= 2

		//decrypt
		plaintext := scheme.Decrypt(encryptedValue)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 10)
		t.Log(plaintext)
	})
	t.Run("encrypt-add-two-cipher-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		a := scheme.Encrypt(20)
		assert.NotNil(t, a)
		assert.NotEqual(t, a, 20)
		t.Log(a)

		b := scheme.Encrypt(30)
		assert.NotNil(t, b)
		assert.NotEqual(t, b, 30)
		t.Log(b)

		c := a + b

		//decrypt
		plaintext := scheme.Decrypt(c)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 50)
		t.Log(plaintext)
	})
	t.Run("encrypt-sub-two-cipher-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)

		//encrypt
		a := scheme.Encrypt(50)
		assert.NotNil(t, a)
		assert.NotEqual(t, a, 50)
		t.Log(a)

		b := scheme.Encrypt(10)
		assert.NotNil(t, b)
		assert.NotEqual(t, b, 10)
		t.Log(b)

		c := a - b

		//decrypt
		plaintext := scheme.Decrypt(c)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 40)
		t.Log(plaintext)
	})
	t.Run("encrypt-mul-two-cipher-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewIntegerToyHomoScheme()
		err := scheme.Generate()
		assert.NoError(t, err)
		t.Log(scheme)

		//encrypt
		a := scheme.Encrypt(20)
		assert.NotNil(t, a)
		assert.NotEqual(t, a, 20)

		b := scheme.Encrypt(30)
		assert.NotNil(t, b)
		assert.NotEqual(t, b, 30)

		c := a * b

		//decrypt
		t.Log(a, b, c)

		plaintext := scheme.Decrypt(c)
		assert.NotNil(t, plaintext)
		assert.Equal(t, plaintext, 600)
		t.Log(plaintext)
	})
}
