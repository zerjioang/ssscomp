from ctypes import cdll

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
    info_python()
    # ml_kit_version()

def main():
    info()
    lib = cdll.LoadLibrary('./s3go.so')
    print "loaded Secret Sharing Library .so file"
    result = lib.add(2, 3)
    print result

if __name__ == '__main__':
    main()