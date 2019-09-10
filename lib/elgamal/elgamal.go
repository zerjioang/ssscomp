// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package elgamal implements ElGamal encryption, suitable for OpenPGP,
// as specified in "A Public-Key Cryptosystem and a Signature Scheme Based on
// Discrete Logarithms," IEEE Transactions on Information Theory, v. IT-31,
// n. 4, 1985, pp. 469-472.
//
// This form of ElGamal embeds PKCS#1 v1.5 padding, which may make it
// unsuitable for other protocols. RSA should be used in preference in any
// case.
package elgamal

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"github.com/zerjioang/ssscomp/lib/common"
	"io"
	"math/big"
)

const (
	paddingPSSize = 11
)

var (
	ErrMessageTooLong = errors.New("elgamal: message too long")
)

// PublicKey represents an ElGamal public key.
type PublicKey struct {
	G, P, Y *big.Int
}

// PrivateKey represents an ElGamal private key.
type PrivateKey struct {
	PublicKey
	X *big.Int
}

// EncryptPadded encrypts the given message to the given public key. The result is a
// pair of integers. Errors can result from reading random, or because msg is
// too large to be encrypted to the public key.
func (pk *PublicKey) EncryptPadded(random io.Reader, msg []byte) (*Cypher, error) {
	pkl := pk.P.BitLen()
	pLen := (pkl + 7) / 8
	if len(msg) > pLen-paddingPSSize {
		err := ErrMessageTooLong
		return nil, err
	}
	// EM = 0x02 || PS || 0x00 || M
	em := make([]byte, pLen-1)
	em[0] = 2
	ps, mm := em[1 : len(em)-len(msg)-1], em[len(em)-len(msg):]
	err := common.FillNonZeroRandomBytes(ps, random)
	if err != nil {
		return nil, err
	}
	em[len(em)-len(msg)-1] = 0
	copy(mm, msg)
	m := new(big.Int).SetBytes(em)

	k, err := rand.Int(random, pk.P)
	if err != nil {
		return nil, err
	}

	c1 := new(big.Int).Exp(pk.G, k, pk.P)
	s := new(big.Int).Exp(pk.Y, k, pk.P)
	c2 := s.Mul(s, m)
	c2.Mod(c2, pk.P)

	c := new(Cypher)
	c.C1 = c1
	c.C2 = c2
	return c, nil
}

func (pk *PublicKey) EncryptNoPadding(random io.Reader, msg []byte) (cypher *Cypher, err error) {
	m := new(big.Int).SetBytes(msg)

	k, err := rand.Int(random, pk.P)
	if err != nil {
		return
	}

	c1 := new(big.Int).Exp(pk.G, k, pk.P)
	s := new(big.Int).Exp(pk.Y, k, pk.P)

	c2 := s.Mul(s, m)
	c2.Mod(c2, pk.P)
	cypher = &Cypher{c1, c2, pk.P}
	return
}

// Encodes public key as hexadecimal JSON
func (pk *PublicKey) GetJson() ([]byte, error) {
	return json.Marshal(map[string]string{
		"g": common.BigIntAsHex(pk.G),
		"pk": common.BigIntAsHex(pk.P),
		"y": common.BigIntAsHex(pk.Y),
	})
}

// DecryptPadded takes two integers, resulting from an ElGamal encryption, and
// returns the plaintext of the message. An error can result only if the
// ciphertext is invalid. Users should keep in mind that this is a padding
// oracle and thus, if exposed to an adaptive chosen ciphertext attack, can
// be used to break the cryptosystem.  See ``Chosen Ciphertext Attacks
// Against Protocols Based on the RSA Encryption Standard PKCS #1'', Daniel
// Bleichenbacher, Advances in Cryptology (Crypto '98),
func (sk *PrivateKey) DecryptPadded(cypher *Cypher) (msg []byte, err error) {
	s := new(big.Int).Exp(cypher.C1, sk.X, sk.P)
	s.ModInverse(s, sk.P)
	s.Mul(s, cypher.C2)
	s.Mod(s, sk.P)
	em := s.Bytes()

	firstByteIsTwo := subtle.ConstantTimeByteEq(em[0], 2)

	// The remainder of the plaintext must be a string of non-zero random
	// octets, followed by a 0, followed by the message.
	//   lookingForIndex: 1 iff we are still looking for the zero.
	//   index: the offset of the first zero byte.
	var lookingForIndex, index int
	lookingForIndex = 1

	for i := 1; i < len(em); i++ {
		equals0 := subtle.ConstantTimeByteEq(em[i], 0)
		index = subtle.ConstantTimeSelect(lookingForIndex&equals0, i, index)
		lookingForIndex = subtle.ConstantTimeSelect(equals0, 0, lookingForIndex)
	}

	if firstByteIsTwo != 1 || lookingForIndex != 0 || index < 9 {
		return nil, errors.New("elgamal: decryption error")
	}
	return em[index+1:], nil
}

func (sk *PrivateKey) DecryptNoPadding(cypher *Cypher) (msg []byte, err error) {
	s := new(big.Int).Exp(cypher.C1, sk.X, sk.P)
	s.ModInverse(s, sk.P)
	s.Mul(s, cypher.C2)
	s.Mod(s, sk.P)
	em := s.Bytes()
	return em, nil
}