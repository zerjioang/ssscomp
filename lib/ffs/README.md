# Feige-Fiat-Shamir

With Feige-Fiat-Shamir zero-knowledge proof we can prove that we know something with actually revealing our information.
The Feige–Fiat–Shamir identification scheme, however, uses modular arithmetic and a parallel verification process that limits the number of communications between Peggy and Victor. 

## Assumptions

* Strong RSA factorization

## Features

* Interactive protocol
* Zero Knowledge Proof

## Setup

Choose two large prime integers p and q and compute the product

```
n = pq
```

Create secret numbers

```
s1,...,sk
```
coprime to n.

Compute v_i ≡ s_i^2 ( mod n ).
Peggy and Victor both receive `n` while `p` and `q` are kept secret.
Peggy is then sent the numbers s_i. These are her secret login numbers.
Victor is sent the numbers v i v_i by Peggy when she wishes to identify herself to Victor.
Victor is unable to recover Peggy's s_i numbers from his v_i numbers due to the difficulty in determining a modular square root when the modulus' factorization is unknown. 

## Security

In the procedure, Peggy does not give any useful information to Victor. She merely proves to Victor that she has the secret numbers without revealing what those numbers are. Anyone who intercepts the communication between each Peggy and Victor would only learn the same information. The eavesdropper would not learn anything useful about Peggy's secret numbers.

## References

* https://link.springer.com/content/pdf/10.1007/3-540-47721-7_12.pdf
* https://asecuritysite.com/encryption/z2
* https://en.wikipedia.org/wiki/Feige%E2%80%93Fiat%E2%80%93Shamir_identification_scheme