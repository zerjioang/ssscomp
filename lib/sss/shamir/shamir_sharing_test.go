package shamir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShamirSharingScheme(t *testing.T) {
	t.Run("3-participant-homomorphic-addition", func(t *testing.T) {
		scheme := NewShamirSharingScheme(1234, 6, 3)
		assert.NotNil(t, scheme)
		assert.Equal(t, scheme.N, 6)
		assert.Equal(t, scheme.secret, 1234)

		// now, lets split our secret so that
		// We wish to divide the secret into 6 parts ( n = 6 )
		// where any subset of 3 parts ( k = 3 ) is sufficient to reconstruct the secret.
		// At random we obtain k − 1 166 and 94.
		scheme.Generate()
		scheme.Print()
	})
}
