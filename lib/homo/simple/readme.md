# Simple Homomorphic Cipher

This is a simple Homomorphic Cipher and involves an odd integer (p) and two random numbers (q and r).
We encrypt a single bit (m) with (p×q)+2×r+m.

Then we take mod p, and then mod 2, and we get the message back

## Example

```bash
M:		1
P:		1001 (odd number)
q:		13 (random number)
r:		9 (random number)

---------------------
Cipher:		13032
Decipher:	1
Decipher matches message
```

## Theory

This is one of the simplest ciphers, where we can add and multiply the ciphers. First we get an odd number (p) and two random numbers (q and r), and encrypt a single bit (m) with:

```math
c=p×q+2×r+m
```

If 2r is smaller than p/2, we cipher with mod p we get:

```math
d=(cmodp)mod2
```

We have basically added noise to the value with the q and r values.

# References
* https://asecuritysite.com/encryption/hom_mill