#  Unpadded RSA
## Homomorphic Secure & Privacy Preserving computation

We show that the Unpadded RSA encryption scheme can be used a Partially Homomorphic Encryption scheme supporting multiplication in the encrypted domain.
Given an RSA public key `(n,e)` and two plaintext `x1` and `x2` we observe:

\begin{equation}
Enc(x1) * Enc(x2) = x1^e * x2^e mod n
                  = (x1^e * x2^e) mod n
                  = Enc(x1^e * x2^e) mod n
\end{equation}

Hence, to compute multiplication in the encrypted domain for RSA, we just need to multiply the ciphertexts to obtain our new ciphertext.

### Example

Following example shows how to ciphertext are homomorphically multiplied using Golang

```go
func ExampleHomomorphicRsaMul(){
	priv, pub := GenerateKeyPair(1024)
	b1 := big.NewInt(int64(200))
	b2 := big.NewInt(int64(3))

	cipher1, _ := EncryptUnpaddedRSA(pub, b1)
	fmt.Println("encrypted x1: ", BigIntAsHex(cipher1))

	cipher2, _ := EncryptUnpaddedRSA(pub, b2)
	fmt.Println("encrypted x2: ", BigIntAsHex(cipher2))

	c12 := big.NewInt(int64(0))
	c12.Mul(cipher1, cipher2)

	decrypted12, _ := DecryptUnpaddedRSA(priv, c12)
	plainMul := decrypted12.String()
	fmt.Println("decrypted value: ", plainMul)
}
```