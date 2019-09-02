from ctypes import *

import wraps
from loader import load_library


# singleton decorator definition
from model.context import SecureContext, ContextType


def singleton(class_):
    instances = {}

    def getinstance(*args, **kwargs):
        if class_ not in instances:
            instances[class_] = class_(*args, **kwargs)
        return instances[class_]

    return getinstance


@singleton
class SSSComp(object):
    # reference to shared object library
    shared_object = None

    def __init__(self):
        print("new sss instance created")
        self.shared_object = load_library()

    def version(self) -> float:
        """
        returns the actual version of the library
        """
        self.shared_object.version.restype = c_double
        current_version = self.shared_object.version()
        return current_version

    def new_smpc_additive(self, participants: int) -> SecureContext:
        """
        Creates a SMPC context based on additive secret sharing schema
        :param participants: number of participants involved in current context
        :return: the context configured
        """
        ctx = SecureContext()
        ctx.shares = participants
        ctx.mode = ContextType.SMPC_ADDITIVE_SHARING
        # call shared library
        self.shared_object.new_smpc_additive.argtypes = [c_int]
        self.shared_object.new_smpc_additive.restype = wraps.SMPCAdditive
        result = self.shared_object.new_smpc_additive(participants)
        print(result)
        return ctx

    def new_smpc_shamir(self, participants: int, minimum: int) -> SecureContext:
        """
        Creates a SMPC context based on shamir secret sharing schema
        :param participants: number of participants (total shares) involved in current context
        :param minimum: minimum number of responses (shares) required to reconstruct encrypted message
        :return: the context configured
        """
        ctx = SecureContext()
        ctx.shares = participants
        ctx.minimum = minimum
        ctx.mode = ContextType.SMPC_SHAMIR_SHARING
        # call shared library
        self.shared_object.new_smpc_shamir(participants, minimum)
        return ctx
