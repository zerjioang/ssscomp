# Fully homomorphic encryption in ring of binary integers

The scheme of fully homomorphic encryption which was proposed by Craig Gentry can be considered in detail by the example calculations in PGSH Z2, and it is analogy for working with bits.

## Description

Firstly, choose any odd **p = 2k + 1** which is secret parameter. Assume **m** in {0, 1}

### Encryption
Encryption function is each **m** is associated with **c = z + pq** where **q** is an arbitrary value. This implies **c = 2r + m + (2k + 1) * q.  Easy to see, that **c mod 2 = (m + q) mod 2** therefore, breaker is able only to determine a parity of encryption result.

### Decryption

Let we know **c** — encrypted number and **p** — known secret. Then the decryption process is: 