# Edwards25519 Elliptic Curve

This is the standard edwards25519 curve definition.

## Introduction

A drop-in replacement for `golang/crypto/ed25519` ([godoc](https://godoc.org/golang.org/x/crypto/ed25519),
[github](https://github.com/golang/crypto/tree/master/ed25519))
 with additional functionality.

# Motivation
In order to verify the validity of a given signature, the validator should posses the public key of the signer. It can be sent along with the message and its signature, which means that the overall data being sent includes 256 bits of the public key. Our function allows to extract the public key from the signature (and the message), thus the public key may not be sent, resulting in a smaller transferred data. Note: there's a computational cost for extracting the public key, so one should consider the trade-off between computations and data size.

## Curve equation

```bash
−x^2 + y^2 = 1 − (121665/121666) * x^2 * y^2
```

Note:

* curve is in two dimensions (nothing fancy, like all the curves is high school)
* curve is mirrored below y axis due to y^2 part of the equation (not a polynomial)

## Base point: G

The base point is a specific point on the curve. It is used as a basis for further calculations. It is an arbitrary choice by the curve authors, just to standardize the scheme.

Note that it is enough to specify the y value and the sign of the x value. That's because the specific x can be calculated from the curve equation.

```bash
G = (x, 4/5)  # take the point with the positive x

# The hex representation of the base point
5866666666666666666666666666666666666666666666666666666666666666
```

## Prime order of the base point: `l`

```bash
l = 2^252 + 27742317777372353535851937790883648493
# =&gt; 7237005577332262213973186563042994240857116359379907606001950938285454250989
```

The `l` is a prime number specified by the curve authors.

In practice this is the private key's strength.

### Total number of points on the curve¶

The total number of points on the curve is also a prime number:

```bash
q = 2^255 - 19
```

In practice not all points are "useful" and so the private key strength is limited to `l` describe above.

## API

### Sign2
Sign2 signs the message with privateKey and returns a signature.
The signature may be verified using Verify2(), if the signer's public key is known.
The signature returned by this method can be used together with the message
to extract the public key using ExtractPublicKey()
It will panic if len(privateKey) is not PrivateKeySize.

```
func Sign2(privateKey PrivateKey, message []byte) []byte
```

### ExtractPublicKey
ExtractPublicKey extracts the signer's public key given a message and its signature.
It will panic if len(sig) is not SignatureSize.

```
func ExtractPublicKey(message, sig []byte) PublicKey
```

### Verify2
Verify2 verifies a signature created with Sign2(), assuming the verifier possesses the public key.

```
func Verify2(publicKey PublicKey, message, sig []byte) bool
````

## Building
```
go build
```

## Testing
```
go test ./... -v
```

## Benchmarking
```
go test -bench=.
go test -bench . github.com/spacemeshos/ed25519/internal/edwards25519
```
