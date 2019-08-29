package paillier

import "math/big"

type PublicKey1 struct {
	G, P, Y *big.Int
}

type PrivateKey1 struct {
	PublicKey1
	X *big.Int
}
