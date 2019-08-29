import os
import sys
from ctypes import *

# shared library version
shared_lib_version = 0.1
notfound_msg = "secure_computation_module not found"


def info_python():
    # Python version
    import sys
    print('Python: {}'.format(sys.version))


def ml_kit_version():
    # scipy
    import scipy
    print('scipy: {}'.format(scipy.__version__))
    # numpy
    import numpy
    print('numpy: {}'.format(numpy.__version__))
    # matplotlib
    import matplotlib
    print('matplotlib: {}'.format(matplotlib.__version__))
    # pandas
    import pandas
    print('pandas: {}'.format(pandas.__version__))
    # scikit-learn
    import sklearn
    print('sklearn: {}'.format(sklearn.__version__))


def info():
    print("Loading Secret Sharing & Secure Computation Library")
    info_python()
    # ml_kit_version()


def load_library():
    info()
    dirname = os.path.dirname(__file__)
    so_name = '{dn}/../../lib/wrapper/secure_computation_ssscomp_{version}.so'
    shared_object = so_name.format(dn=dirname, version=shared_lib_version)
    print("so file = ", shared_object)
    try:
        lib = cdll.LoadLibrary(shared_object)
        if lib is not None:
            # 1 print library version
            lib.version.restype = c_double
            current_version = lib.version()
            print("version = ", current_version)

            # 2 print library banner
            print("loaded secure_computation_module from .so file")
            lib.hello()

            # run example function to check compatibility
            lib.Add.argtypes = [c_longlong, c_longlong]
            # lib.Cosine.restype = c_double
            print("lib.Add(12,99) = %d" % lib.Add(12, 99))
            return lib
    except FileNotFoundError:
        print(notfound_msg)
        secure_computation_module = None
        print("fatal error: secure_computation_module not found", file=sys.stderr)
        exit(127)
    except ModuleNotFoundError:
        print(notfound_msg)
        secure_computation_module = None
        print("fatal error: secure_computation_module not found", file=sys.stderr)
        exit(127)
    except Exception:
        secure_computation_module = None
        print("fatal error: secure_computation_module not found", file=sys.stderr)
        exit(127)
    return None
