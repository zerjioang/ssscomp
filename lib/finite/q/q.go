// Package q implements Q field arithmetic operations over
// user specified field usually based on a prime Q, hence the name
package q

import (
	"errors"
	"fmt"
	"math/big"
)

// QField creates a finite field element n over p
type QField struct {
	n big.Int // n is in {0..p-1}
	p big.Int // Fp
}

// NewQField takes an element n and p an order
// Also called Galois fields are a body defined on a
// finite set of elements where the number of elements (n)
// must be prime or a power of a prime number so that
// n = p^m
// If m=1 its a particular case called prime field.
// n = p
// prime fields
// Known as (Galois Field) or GF(p^1) are formed by a set of (p-1) elements and operations of bodies defined
// Example:
// GF(7)
// Elementos: {0,1,2,3,4,5,6}
// Operations: {+, -, *, /}
// 1 + 1 = 2 mod 7 = 2
// 1 + 6 = 7 mod 7 = 0
// 2 + 6 = 8 mod 7 = 1
//
// 2 * 3 = 6 mod 7 = 6
// 3 * 3 = 9 mod 7 = 2
// So far so integer, but what happens when m > 1 ?
//
// Extension fields
// We will call extension field to a finite body where the number of elements
// is a power of a prime number and the elements of the set are polynomials of
// the form:
// am−1xm−1+...+a0x+a0
// https://iagolast.github.io/blog/2016/11/06/campos-galois.html
// http://www.miscelaneamatematica.org/Misc53/5306.pdf
func NewQField(n int64, p int64) (*QField, error) {
	if n >= p || n < 0 {
		return nil, errors.New("n must be a positive integer ")
	}
	bigN := big.NewInt(n)
	bigP := big.NewInt(p)
	Fp := QField{*bigN, *bigP}
	return &Fp, nil
}

// Equal is a method to compare two Field Elements and return a bool
func (f *QField) Equal(e QField) bool {
	return e.n.Cmp(&f.n) == 0 && e.p.Cmp(&f.p) == 0
}

// NotEqual is the opposite of Equal
func (f *QField) NotEqual(e QField) bool {
	return !f.Equal(e)
}

// Print will print the element in a cool way
func (f *QField) Print() {
	fmt.Printf("FieldElement_%s(%s)", f.p.Text(10), f.n.Text(10))
}

// Add will add f and e and return their sum
func Add(e, f QField) (*QField, error) {
	// check that e and f are over the same p
	if e.p.Cmp(&f.p) != 0 {
		return nil, errors.New("both elements should be on the same p")
	}
	sum := big.NewInt(0)
	sum.Add(&e.n, &f.n)
	z := big.NewInt(0)
	z.Mod(sum, &e.p)
	return &QField{*z, e.p}, nil
}

// Sub will substract f from e and return their dff
func Sub(e, f QField) (*QField, error) {
	// check that e and f are over the same p
	if e.p.Cmp(&f.p) != 0 {
		return nil, errors.New("both elements should be on the same p")
	}
	red := big.NewInt(0)
	red.Sub(&e.n, &f.n)
	z := big.NewInt(0)
	z.Mod(red, &e.p)
	return &QField{*z, e.p}, nil
}

// Mul will multiply two field elements
func Mul(e, f QField) (*QField, error) {
	// check that e and f are over the same p
	if e.p.Cmp(&f.p) != 0 {
		return nil, errors.New("both elements should be on the same p")
	}
	mul := big.NewInt(0)
	mul.Mul(&e.n, &f.n)
	z := big.NewInt(0)
	z.Mod(mul, &e.p)
	return &QField{*z, e.p}, nil
}

// Pow will return the power of a field element to a power
func Pow(e QField, power int64) (*QField, error) {
	// power reduction
	//p := power % (e.p - 1)

	bPower := big.NewInt(0)
	bPower.Sub(&e.p, big.NewInt(1))

	bP := big.NewInt(0)
	bP.Mod(big.NewInt(power), bPower)
	// n**(p-1) % p
	//n := math.Pow(float64(e.n), float64(p))

	bN := big.NewInt(0)
	bN.Exp(&e.n, bP, &e.p)

	return &QField{*bN, e.p}, nil

}
