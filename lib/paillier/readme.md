# Partial Homomorphic Encryption with ElGamal (Multiply)

With ElGamal public key encryption we have a public key of (Y,g,p) and a private key of (x). The cipher has two elements (a and b). With this, Bob selects a private key of x and the calculates Y (gx(modp)) for his public key. We can use it for partial homomorphic encrytion, and where we can multiply the ciphered values. With this we encrypt two integers, and then multiply the ciphered values, and then decrypt the result. 

## Other implementations

* https://github.com/actuallyachraf/gomorph

## References

* https://github.com/mortendahl/privateml/blob/master/paillier/Paillier.ipynb