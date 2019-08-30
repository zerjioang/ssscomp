package unpaddedrsa

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKey(t *testing.T) {
	t.Run("generate-prime", func(t *testing.T) {
		primeList, err := GeneratePrime(rand.Reader, 3, 64)
		assert.NoError(t, err)
		assert.NotNil(t, primeList)
		assert.Equal(t, len(primeList), 3)
		t.Log(
			BigIntAsDecimal(primeList[0]),
			BigIntAsDecimal(primeList[1]),
			BigIntAsDecimal(primeList[2]),
		)
	})
}
