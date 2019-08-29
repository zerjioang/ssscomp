from ctypes import *
from loader import load_library


# singleton decorator definition
def singleton(class_):
    instances = {}

    def getinstance(*args, **kwargs):
        if class_ not in instances:
            instances[class_] = class_(*args, **kwargs)
        return instances[class_]

    return getinstance


# our shared_object library wrapper access
@singleton
class SSS(object):
    # reference to shared object library
    shared_object = None

    def __init__(self):
        print("new sss instance created")
        self.shared_object = load_library()
        
    def version(self) -> float:
        self.shared_object.version.restype = c_double
        current_version = self.shared_object.version()
        return current_version
