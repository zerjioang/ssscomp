package paillier

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

// This example demonstrates basic usage of this library.
// Features shown:
//   * Encrypt/Decrypt
//   * Homomorphic cipher text addition
//   * Homomorphic addition with constant
//   * Homomorphic multiplication with constant
func TestFullPaillier(t *testing.T) {
	t.Run("128-bits-key", func(t *testing.T) {
		runPaillier(128)
	})
	t.Run("192-bits-key", func(t *testing.T) {
		runPaillier(192)
	})
	t.Run("256-bits-key", func(t *testing.T) {
		runPaillier(256)
	})
	t.Run("512-bits-key", func(t *testing.T) {
		runPaillier(512)
	})
	t.Run("1024-bits-key", func(t *testing.T) {
		runPaillier(1024)
	})
	t.Run("2048-bits-key", func(t *testing.T) {
		runPaillier(2048)
	})
	t.Run("4096-bits-key", func(t *testing.T) {
		runPaillier(4096)
	})
}

func runPaillier(keySize int) {
	// Encrypt a 128-bit private key.
	privKey, err := GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encrypt the number "15".
	m15 := new(big.Int).SetInt64(15)
	c15, err := Encrypt(&privKey.PublicKey2, m15.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decrypt the number "15".
	d, err := Decrypt(privKey, c15)
	if err != nil {
		fmt.Println(err)
		return
	}
	plainText := new(big.Int).SetBytes(d)
	fmt.Println("Decryption Result of 15: ", plainText.String())

	// Now for the fun stuff.
	// Encrypt the number "20".
	m20 := new(big.Int).SetInt64(20)
	c20, err := Encrypt(&privKey.PublicKey2, m20.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add the encrypted integers 15 and 20 together.
	plusM16M20 := AddCipher(&privKey.PublicKey2, c15, c20)
	decryptedAddition, err := Decrypt(privKey, plusM16M20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result of 15+20 after decryption: ", new(big.Int).SetBytes(decryptedAddition).String()) // 35

	// Add the encrypted integer 15 to plaintext constant 10.
	plusE15and10 := Add(&privKey.PublicKey2, c15, new(big.Int).SetInt64(10).Bytes())
	decryptedAddition, err = Decrypt(privKey, plusE15and10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result of 15+10 after decryption: ", new(big.Int).SetBytes(decryptedAddition).String()) // 25

	// Multiply the encrypted integer 15 by the plaintext constant 10.
	mulE15and10 := Mul(&privKey.PublicKey2, c15, new(big.Int).SetInt64(10).Bytes())
	decryptedMul, err := Decrypt(privKey, mulE15and10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result of 15*10 after decryption: ", new(big.Int).SetBytes(decryptedMul).String()) // 150
}
