## Finite field or Galois field

A Galois field, also known as a finite field, is a mathematical field with a number of elements equal to a prime number to a positive integer power. While finite fields with a prime number of elements are familiar to most programmers -- boolean arithmetic, a.k.a. arithmetic mod 2 is a well-known example -- fields that take that prime to powers higher than 1 are less well-known.

Basically, an element of GF(2**m) can be seen as a list of m bits, where addition and multiplication are elementwise mod 2 (a XOR b for addition, a AND b for multiplication) and the remaining rules of field arithmetic follow from linear algebra (vectors, or alternatively, polynomial coefficients).

Short version: an element of GF(2**8) element may be represented as a byte (0 ≤ n ≤ 255), but it's really a vector of 8 bits -- like a very primitive MMX/SSE. We then treat said vector as the coefficients of a polynomial, and that allows us to define multiplication, giving us a full mathematical field.

Finite fields -- and GF(2**8) in particular -- get a ton of use in codes, in both the "error-correcting code" and "cryptographic code" senses. However, this implementation has NOT been hardened against timing attacks, so it MUST NOT be used in cryptography.

## Computing fields

Get instance of your desired GF(x)

```go
gf, err := galois.GF(8)
```

And make modulu based computations like:

```go
mul, err := gf.Mul(3, 4)
div, err := gf.Div(6, 43)
expon, err := gf.Expon(34, 9001)
```
## Other implementations

* https://github.com/betawaffle/gf256
* https://github.com/cloud9-tools/go-galoisfield
