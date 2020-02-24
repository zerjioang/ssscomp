package bulletproof

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
)

// serializedPedersenCommitment is the constant that is encoded to signal that
// the encoded value is a Pedersen commitment, rather than a standard compressed
// curve point.
const serializedPedersenCommitment byte = 0x9

var (
	// describe a (0,0) EC point
	pZero = ECPoint{big.NewInt(0), big.NewInt(0)}
)

// ECPoint is a group element of the secp256k1 curve in affine coordinates.
type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// Equals returns true if the given point is the same.
func (p *ECPoint) Equals(other *ECPoint) bool {
	return p != nil &&
		other != nil &&
		p.X.Cmp(other.X) == 0 &&
		p.Y.Cmp(other.Y) == 0
}

// Mult multiplies point p by scalar s and returns the resulting point
func (p *ECPoint) Mult(s *big.Int) ECPoint {
	modS := new(big.Int).Mod(s, curve.N)
	X, Y := curve.ScalarMult(p.X, p.Y, modS.Bytes())
	return ECPoint{X, Y}
}

// Add adds points p and p2 and returns the resulting point
func (p *ECPoint) Add(p2 ECPoint) ECPoint {
	X, Y := curve.Add(p.X, p.Y, p2.X, p2.Y)
	return ECPoint{X, Y}
}

// Neg returns the additive inverse of point p
func (p *ECPoint) Neg() ECPoint {
	negY := new(big.Int).Neg(p.Y)
	modValue := negY.Mod(negY, curve.Params().P) // mod P is fine here because we're describing a curve point
	return ECPoint{p.X, modValue}
}

// String prints the coordinates of this point.
func (p *ECPoint) String() string {
	return fmt.Sprintf("{x: %032x, y: %032x}", p.X.Bytes(), p.Y.Bytes())
}

// Read deserializes a compressed elliptic curve point from the reader.
func (p *ECPoint) Read(r io.Reader) error {
	buf := make([]byte, 32+1)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	sign := buf[0]
	x := new(big.Int).SetBytes(buf[1:])

	if (sign & 0xfe) != 8 {
		return errors.New("point is not serialized correctly")
	}

	// Derive the possible y coordinates from the secp256k1 curve
	// y² = x³ + 7.
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)
	x3.Add(x3, curve.Params().B)

	// y = ±sqrt(x³ + 7).
	y := ModSqrtFast(x3)

	// Pick which y from the sign encoded in the first byte.
	if (sign & 1) != 0 {
		y = new(big.Int).Sub(curve.P, y)
	}

	p.X = x
	p.Y = y

	return nil
}

/*
Vector Pedersen Commitment
Given an array of values, we commit the array with different generators
for each element and for each randomness.

Bytes compresses and serializes the point.
*/
func (p *ECPoint) Bytes() []byte {
	buff := new(bytes.Buffer)
	sign := serializedPedersenCommitment
	if IsQuadraticResidue(p.Y) {
		sign ^= 1
	}
	if err := buff.WriteByte(sign); err != nil {
		log.Fatal(err)
	}
	x := GetB32(p.X)
	if _, err := buff.Write(x[:]); err != nil {
		log.Fatal(err)
	}
	return buff.Bytes()
}
