#ElGamal

In cryptography, the ElGamal encryption system is an asymmetric key encryption algorithm for public-key cryptography which is based on the 
Diffie–Hellman key exchange. It was described by Taher Elgamal in 1985. ElGamal encryption is used in the free GNU Privacy Guard software,
recent versions of PGP, and other cryptosystems. The Digital Signature Algorithm (DSA) is a variant of the ElGamal signature scheme, 
which should not be confused with ElGamal encryption.

ElGamal encryption can be defined over any cyclic group `G`, such as multiplicative group of integers modulo n. Its security depends upon 
the difficulty of a certain problem in `G` related to computing discrete logarithms. 

## Description

The algorithm
ElGamal encryption consists of three components:

* the key generator
* the encryption algorithm
* the decryption algorithm.

The ElGamal paper and the Handbook of Applied Cryptography state to select the private key `x` in the range `{ 1 , … , p − 1 }`
Tsiounis and Yung showed the lower limit as `{ 1 , … , q − 1 }` in On the Security of ElGamal Based Encryption