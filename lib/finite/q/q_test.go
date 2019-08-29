// testing
package q

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNewQField(t *testing.T) {
	r, err := NewQField(2, 31)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	caseN := big.NewInt(2)
	caseP := big.NewInt(31)
	if r.n.Cmp(caseN) != 0 && r.p.Cmp(caseP) != 0 || err != nil {
		t.Fatalf("Expected %d,%d but got %v", caseN, caseP, r)
	}
}

func TestQFieldEqual(t *testing.T) {
	a := QField{*big.NewInt(15), *big.NewInt(31)}
	b := QField{*big.NewInt(15), *big.NewInt(31)}
	result := a.Equal(b)
	if result != true {
		t.Fatalf("Expected true, but got %v", result)
	}
}

func TestQFieldNotEqual(t *testing.T) {
	a := QField{*big.NewInt(15), *big.NewInt(31)}
	b := QField{*big.NewInt(7), *big.NewInt(31)}
	result := a.NotEqual(b)
	if result != true {
		t.Fatalf("Expected true, but got %v", result)
	}
}

func TestQFieldAdd(t *testing.T) {

	a := QField{*big.NewInt(17), *big.NewInt(31)}
	b := QField{*big.NewInt(21), *big.NewInt(31)}
	result, err := Add(a, b)
	if err != nil {
		t.Fatalf("Error : %v", err)
	}
	truth := QField{*big.NewInt(7), *big.NewInt(31)}
	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}
}

func TestQFieldSub(t *testing.T) {
	a := QField{*big.NewInt(29), *big.NewInt(31)}
	b := QField{*big.NewInt(4), *big.NewInt(31)}
	result, err := Sub(a, b)
	if err != nil {
		t.Fatalf("Error : %v", err)
	}
	truth := QField{*big.NewInt(25), *big.NewInt(31)}
	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}
}

func TestQFieldMul(t *testing.T) {
	a := QField{*big.NewInt(24), *big.NewInt(31)}
	b := QField{*big.NewInt(19), *big.NewInt(31)}
	result, err := Mul(a, b)
	if err != nil {
		t.Fatalf("Error : %v", err)
	}
	truth := QField{*big.NewInt(22), *big.NewInt(31)}
	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}
}

func TestQFieldPow(t *testing.T) {
	a := QField{*big.NewInt(17), *big.NewInt(31)}
	result, err := Pow(a, 3)
	if err != nil {
		t.Fatalf("Error : %v", err)
	}
	truth := QField{*big.NewInt(15), *big.NewInt(31)}
	if result.NotEqual(truth) {
		t.Fatalf("Expected %v, but got %v", truth, result)
	}

}
