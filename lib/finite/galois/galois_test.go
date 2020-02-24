package galois

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testpoly2 = GfPoly{
	base:   2,
	NW:     1 << 2,
	gflog:  []uint32{0, 3, 1, 2},
	gfilog: []uint32{1, 2, 3, 1},
}

// TestGF tests GF(x)
func TestGF(t *testing.T) {
	t.Run("gf-0", func(t *testing.T) {
		// test 0
		gfpoly, err := GF(0)
		assert.NoError(t, err)
		assert.NotNil(t, gfpoly)
	})
	t.Run("gf-1", func(t *testing.T) {
		var gf uint8
		// test 1
		gfpoly, err := GF(1)
		assert.NoError(t, err)
		assert.NotNil(t, gfpoly)

		// test overflow
		gf = MaxComputableGaloisField + 1
		gfpoly, err = GF(gf)
		assert.NoError(t, err)
		// test limit
		gf = MaxComputableGaloisField
		gfpoly, err = GF(gf)
		assert.NoError(t, err)
	})
	t.Run("gf-2", func(t *testing.T) {
		gfpoly, err := GF(2)
		assert.NoError(t, err)
		assert.NotNil(t, gfpoly)
		// test 2
		if !reflect.DeepEqual(&testpoly2, gfpoly) {
			fmt.Printf("testpoly2: %+v\n", testpoly2)
			fmt.Printf("gfpoly: %+v\n", gfpoly)
			t.Error("GF(", 2, ") is not good :c")
		}
	})
	t.Run("gf-2-operations", func(t *testing.T) {
		gfpoly, err := GF(2)
		assert.NoError(t, err)
		assert.NotNil(t, gfpoly)
		result, err := gfpoly.Mul(111, 101)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, 11011)
	})
}

//TestMul tests mul
func TestMul(t *testing.T) {

	var cal uint32

	var testmul2 = []struct {
		test   []uint32
		result uint32
	}{
		{[]uint32{0, 0}, 0},
		{[]uint32{0, 1}, 0},
		{[]uint32{0, 2}, 0},
		{[]uint32{0, 3}, 0},
		{[]uint32{1, 0}, 0},
		{[]uint32{1, 1}, 1},
		{[]uint32{1, 2}, 2},
		{[]uint32{1, 3}, 3},
		{[]uint32{2, 0}, 0},
		{[]uint32{2, 1}, 2},
		{[]uint32{2, 2}, 3},
		{[]uint32{2, 3}, 1},
		{[]uint32{3, 0}, 0},
		{[]uint32{3, 1}, 3},
		{[]uint32{3, 2}, 1},
		{[]uint32{3, 3}, 2},
	}

	gfpoly, err := GF(2)
	if err != nil {
		t.Error(err)
	}
	// test in range
	for _, mul := range testmul2 {
		cal, err = gfpoly.Mul(mul.test[0], mul.test[1])
		if err != nil {
			t.Error("Mul", mul.test[0], mul.test[1], "fails")
		}
		if cal != mul.result {
			t.Error("Mul", mul.test[0], mul.test[1], " result does not match. Got", cal, "expect", mul.result)
		}
	}

	// test out range
	cal, err = gfpoly.Mul(3, 4)
	if err == nil {
		t.Error("Out of range should fails")
	}
	cal, err = gfpoly.Mul(4, 3)
	if err == nil {
		t.Error("Out of range should fails")
	}
}

//TestMul tests div
func TestDiv(t *testing.T) {

	var cal uint32

	var testdiv2 = []struct {
		test   []uint32
		result uint32
	}{
		{[]uint32{0, 0}, 0},
		{[]uint32{0, 1}, 0},
		{[]uint32{0, 2}, 0},
		{[]uint32{0, 3}, 0},
		{[]uint32{1, 0}, 0},
		{[]uint32{1, 1}, 1},
		{[]uint32{1, 2}, 3},
		{[]uint32{1, 3}, 2},
		{[]uint32{2, 0}, 0},
		{[]uint32{2, 1}, 2},
		{[]uint32{2, 2}, 1},
		{[]uint32{2, 3}, 3},
		{[]uint32{3, 0}, 0},
		{[]uint32{3, 1}, 3},
		{[]uint32{3, 2}, 2},
		{[]uint32{3, 3}, 1},
	}

	gfpoly, err := GF(2)
	if err != nil {
		t.Error(err)
	}
	// test in range
	for _, div := range testdiv2 {
		cal, err = gfpoly.Div(div.test[0], div.test[1])
		if div.test[1] > 0 && err != nil {
			t.Error("Div", div.test[0], div.test[1], "fails")
		}
		if div.test[1] == 0 && err == nil {
			t.Error("Division by zero did not rise error")
		}
		if div.test[1] > 0 && cal != div.result {
			t.Error("Div", div.test[0], div.test[1], " result does not match. Got", cal, "expect", div.result)
		}
	}

	// test out range
	cal, err = gfpoly.Div(3, 4)
	if err == nil {
		t.Error("Out of range should fails")
	}
	cal, err = gfpoly.Div(4, 3)
	if err == nil {
		t.Error("Out of range should fails")
	}
}

//TestExpon tests expon
func TestExpon(t *testing.T) {

	var cal uint32

	var testexpon2 = []struct {
		test   []uint32
		result uint32
	}{
		{[]uint32{0, 0}, 1},
		{[]uint32{0, 1}, 0},
		{[]uint32{0, 2}, 0},
		{[]uint32{0, 3}, 0},
		{[]uint32{1, 0}, 1},
		{[]uint32{1, 1}, 1},
		{[]uint32{1, 2}, 1},
		{[]uint32{1, 3}, 1},
		{[]uint32{2, 0}, 1},
		{[]uint32{2, 1}, 2},
		{[]uint32{2, 2}, 3},
		{[]uint32{2, 3}, 1},
		{[]uint32{3, 0}, 1},
		{[]uint32{3, 1}, 3},
		{[]uint32{3, 2}, 2},
		{[]uint32{3, 3}, 1},
	}

	gfpoly, err := GF(2)
	if err != nil {
		t.Error(err)
	}
	// test in range
	for _, expon := range testexpon2 {
		cal, err = gfpoly.Expon(expon.test[0], expon.test[1])
		if err != nil {
			t.Error("Expon", expon.test[0], expon.test[1], "fails")
		}
		if cal != expon.result {
			t.Error("Expon", expon.test[0], expon.test[1], " result does not match. Got", cal, "expect", expon.result)
		}
	}

	// test out range
	cal, err = gfpoly.Expon(3, 4)
	if err != nil {
		t.Error("3, 4 shouldn't fail")
	}
	cal, err = gfpoly.Expon(4, 3)
	if err == nil {
		t.Error("4, 3 Out of range should fail")
	}
}
