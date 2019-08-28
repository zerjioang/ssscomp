package simple

import "math/rand"

type SimpleHomoCipher struct {
	// odd integer (p)
	p int
	// random number 1
	q int
	// random number 2
	r int
	// message to be encrypted
	m int
	// result
	e int
}

func NewSimpleHomoCipher(v int) SimpleHomoCipher {
	s := SimpleHomoCipher{}
	s.m = v
	return s
}

// set current message to be homomorphically encrypted
func (shc *SimpleHomoCipher) Message(message int) {
	shc.m = message
}

// execute homomorphical encryption with given parameters
func (shc *SimpleHomoCipher) Do() (int, error) {
	shc.p = 1001
	shc.q = rand.Intn(10)
	shc.r = rand.Intn(10)
	shc.e = (shc.p * shc.q) + 2*shc.r + shc.m
	return shc.e, nil
}

// execute homomorphical encryption with given parameters
func (shc *SimpleHomoCipher) Recover(cipherValue int) int {
	return (cipherValue % shc.p) % 2
}

// get string representation of formula
func (shc *SimpleHomoCipher) Describe() string {
	return "( p × q ) + 2 × r + m"
}
