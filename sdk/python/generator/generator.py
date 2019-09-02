import pybindgen
import sys


def gen():
    # Then, we create an object to represent the module we want to generate:
    mod = pybindgen.Module('MyModule')
    # add our C header:
    mod.add_include('"my-module.h"')
    # and register our function which returns no value
    # (hence, the second argument ‘None’), and, takes no arguments
    # (hence, the third argument, the empty list ‘[]’):
    mod.add_function('MyModuleDoAction', None, [])
    # Finally, we generate code for this binding directed to standard output:
    mod.generate(sys.stdout)


if __name__ == '__main__':
    gen()
