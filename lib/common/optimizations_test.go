package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHorner(t *testing.T) {
	r := Horner(3, []int64{-19, 7, -4, 6})
	assert.Equal(t, r, 128)
	t.Log(r)
}
