package fhering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingZtwo(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		a := 5
		b := 3
		t.Log(a, b)

		schema := NewRingZtwo()
		genErr := schema.Generate()

		assert.NoError(t, genErr)
	})
}
