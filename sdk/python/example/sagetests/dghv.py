""" DGHV - Somewhat Homomorphic Encryption Scheme                      """
""" With: key-compression ToDo: bootstrapping&squashed decryption      """
""" 2015 - Gary Cornelius                                              """
from time import time


class DGHV:
    def keyGen(self):
        # Generate secret key
        # Use proof=false for speedup
        set_random_seed(time())
        p = random_prime(2 ^ self.eta, lbound=2 ^ (self.eta - 1), proof=False)
        q_0 = 0

        # Procedure from KeyGen(1^lam)
        # Choose random odd q_0
        while q_0 % 2 == 0:
            q_0 = ZZ.random_element((2 ^ self.gamma) // p)
        x_0 = q_0 * p

        # Set seed for recovery of X_i at encryption
        seed = 0
        set_random_seed(seed)
        # 0 <= X_i < 2^gamma
        X_i = [ZZ.random_element(2 ^ self.gamma) for i in range(self.tao)]

        # Continue normal random
        set_random_seed(time())
        # 0 <= E < 2^(lam+eta)//p
        E = [ZZ.random_element((2 ^ (self.lam + self.eta)) // p) for i in range(self.tao)]
        # -2^rho < r < 2^rho
        r_i = [ZZ.random_element((-2 ^ self.rho) + 1, 2 ^ self.rho) for i in range(self.tao)]
        # Construct d_i
        d_i = [(X_i[i] % p + (E[i] * p) - r_i[i]) for i in range(self.tao)]

        # Return private and secret key pairs
        self.pk = [seed, x_0, d_i]
        self.sk = p

        print
        "keyGen() - done."
        return 1

    def encrypt(self, m):
        seed = self.pk[0]
        x_0 = self.pk[1]
        d_i = self.pk[2]

        # Recover X_i from seed
        set_random_seed(seed)
        X_i = [ZZ.random_element(2 ^ (self.gamma)) for i in range(self.tao)]

        # Generate x_i from X_i and d_i
        x_i = [X_i[i] - d_i[i] for i in range(self.tao)]

        # randomize
        set_random_seed(time())
        b_i = [ZZ.random_element(2 ^ (self.alpha)) for i in range(self.tao)]

        self.rhoo = self.rho + self.alpha  # + w(log(lambda))
        r = ZZ.random_element(-2 ^ (self.rhoo) + 1, 2 ^ (self.rhoo))

        # Initialize cypher
        c = 0
        for i in range(self.tao):
            c = (c + x_i[i] * b_i[i])
        c = (m + 2 * r + 2 * c) % x_0
        return c

    def decrypt(self, c):
        # ToDo: Take care of carry?
        return (c - self.sk * ((c / self.sk).round())) % 2

    def __init__(self, lam, rho, eta, gamma, tao, alpha):
        self.lam = lam
        self.rho = rho
        self.eta = eta
        self.gamma = gamma
        self.tao = tao
        self.alpha = alpha
        self.rhoo = 0
        self.pk = []
        self.sk = 0


def test():
    # global d
    # Parameters from Public Key Compression paper
    d = DGHV(42, 27, 1026, 150000, 158, 200)
    d.keyGen()
    c1 = d.encrypt(1)
    c2 = d.encrypt(0)
    # print "",d.decrypt(c1)
    print
    "Addition"

    print
    "Dec(Enc(1)+Enc(0))=", d.decrypt(c1 + c2)
    print
    "Dec(Enc(0)+Enc(1))=", d.decrypt(c2 + c1)
    print
    "Dec(Enc(1)+Enc(1))=", d.decrypt(c1 + c1)
    print
    "Dec(Enc(0)+Enc(0))=", d.decrypt(c2 + c2)
    print
    "Dec(Enc(0)+Enc(0)+Enc(1))=", d.decrypt(c2 + c2 + c1)
    print
    "Dec(Enc(0)+Enc(0)+Enc(0))=", d.decrypt(c2 + c2 + c2)

    print
    "--------------"
    print
    "Mult"

    print
    "Dec(Enc(1)*Enc(0))=", d.decrypt(c1 * c2)
    print
    "Dec(Enc(0)*Enc(1))=", d.decrypt(c2 * c1)
    print
    "Dec(Enc(1)*Enc(1))=", d.decrypt(c1 * c1)
    print
    "Dec(Enc(0)*Enc(0))=", d.decrypt(c2 * c2)
    print
    "--------------"
    print
    "Mixed"
    print
    "Dec(Enc(1)*Enc(1)+Enc(0)*Enc(1))=", d.decrypt(c1 * c1 + c2 * c1)
