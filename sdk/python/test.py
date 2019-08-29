from lib import SSSComp


def run_test():
    print("executing test program")

    # access to shared module via singleton pattern
    secure_mod = SSSComp()

    # test 1: print current library version
    print("current library version is: ", secure_mod.version())


if __name__ == '__main__':
    run_test()
