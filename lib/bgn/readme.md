# The Boneh Goh Nissim (BGN) Cryptosystem

The  cryptosystem  devised  by  Boneh,  Goh,  and  Nissim  [1]  was  the  first  to  allow  both  additionsand multiplications with a constant-size ciphertext.  There is a catch, however:  while the additiveproperty is the same as for the ElGamal variant, onlyonemultiplication is permitted.  The systemis thus called “somewhat homomorphic.”
One of the key ideas in the BGN system is to use elliptic curve groups whose order is a compositenumbernthat is hard to factor.  (In all previous systems we required the group order to be prime.)

## Limitations

* Homomorphic Addition
* Only 1x Homomorphic Multiplication allowed
