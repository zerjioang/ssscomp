package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShare(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		s := NewIntShare(0)
		assert.NotNil(t, s)
	})
	t.Run("instantiate-ptr", func(t *testing.T) {
		s := NewIntSharePtr(0)
		assert.NotNil(t, s)
	})
	t.Run("string", func(t *testing.T) {
		s := NewIntSharePtr(0)
		str := s.String()
		assert.NotNil(t, str)
		assert.True(t, len(str) > 0)
	})
}
