import random

from lib import SSSComp


def run_test():
    print("executing test program")

    # access to shared module via singleton pattern
    secure_mod = SSSComp()

    # test 1: print current library version
    print("current library version is: ", secure_mod.version())
    print(additive_share(20000, 3392421, 3))


def additive_share(secret, Q, N):
    shares = [random.randrange(Q) for _ in range(N - 1)]
    shares += [(secret - sum(shares)) % Q]
    return shares


if __name__ == '__main__':
    run_test()
