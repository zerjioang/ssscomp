package factoring

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	if fmt.Sprintf("%v", PrimeFactors(23)) != `[23]` {
		t.Error(23)
	}
	if fmt.Sprintf("%v", PrimeFactors(12)) != `[2 2 3]` {
		t.Error(12)
	}
	if fmt.Sprintf("%v", PrimeFactors(360)) != `[2 2 2 3 3 5]` {
		t.Error(360)
	}
	if fmt.Sprintf("%v", PrimeFactors(97)) != `[97]` {
		t.Error(97)
	}
}

func TestPrimeFactorsExamples(t *testing.T) {
	// P = 211, Q = 241, N = 50851
	t.Run("factors-50851", func(t *testing.T) {
		factors := PrimeFactors(50851)
		assert.NotNil(t, factors)
		assert.Equal(t, len(factors), 2)
		assert.Equal(t, factors[0], 211)
		assert.Equal(t, factors[1], 241)
		t.Log(factors)
	})
	// P = 53087, Q = 59497, N = 3158517239
	t.Run("factors-3158517239", func(t *testing.T) {
		factors := PrimeFactors(3158517239)
		assert.NotNil(t, factors)
		assert.Equal(t, len(factors), 2)
		assert.Equal(t, factors[0], 53087)
		assert.Equal(t, factors[1], 59497)
		t.Log(factors)
	})
}
