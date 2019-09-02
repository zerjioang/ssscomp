package toy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigToyHomoScheme(t *testing.T) {
	t.Run("encrypt-decrypt", func(t *testing.T) {
		//generate value 3 as secure value
		scheme := NewBigToyHomoScheme()
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
		assert.Equal(t, plaintext.String(), "3")
		t.Log(plaintext)
	})
}
