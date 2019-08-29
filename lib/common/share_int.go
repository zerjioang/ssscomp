package common

import (
	"fmt"
	"math"
)

// Secret Sharing individual item, also known as IntShare
type IntShare struct {
	Shareable
	v int
}

// Creates new IntShare
func NewIntShare(v int) IntShare {
	return IntShare{v: v}
}

// Creates new IntShare as a ptr
func NewIntSharePtr(v int) *IntShare {
	s := NewIntShare(v)
	return &s
}

// String representation of a IntShare using `fmt` package
func (s *IntShare) String() string {
	return fmt.Sprintf("int share: %d", s.v)
}

func (s *IntShare) IntValue() int {
	return s.v
}

// share operations for homomorphic operations
// operations over current finite field (aka sample value)
func (s *IntShare) Add(b Shareable) (Shareable, error) {
	s.v += b.IntValue()
	return s, nil
}
func (s *IntShare) Sub(b Shareable) (Shareable, error) {
	s.v -= b.IntValue()
	return s, nil
}
func (s *IntShare) Mul(b int) (Shareable, error) {
	s.v *= b
	return s, nil
}
func (s *IntShare) Pow(exponent int) (Shareable, error) {
	r := math.Pow(float64(s.v), float64(exponent))
	s.v = int(r)
	return s, nil
}
func (s *IntShare) Div(b Shareable) (Shareable, error) {
	s.v = s.v / b.IntValue()
	return s, nil
}
func (s *IntShare) Neg() (Shareable, error) {
	s.v = -s.v
	return s, nil
}
func (s *IntShare) Mod(q int) (Shareable, error) {
	s.v = s.v % q
	return s, nil
}
func (s *IntShare) Copy() Shareable {
	return NewIntSharePtr(s.v)
}
func (s *IntShare) Reset() {
	s.v = 0
}
