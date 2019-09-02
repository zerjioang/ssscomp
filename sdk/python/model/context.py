from enum import Enum


class ContextType(Enum):
    """
    List of supported Secure Multi Party Computation secret sharing schemas
    """
    UNDEFINED = 0
    SMPC_ADDITIVE_SHARING = 1
    SMPC_SHAMIR_SHARING = 2


class SecureContext(object):
    """
    Number of participants in current context
    """
    shares = 0
    """
    Number of mimimum participants that needs to respond in order to reconstruct the message
    """
    minimum = 0
    """
    Type of context
    """
    mode = None

    def __init__(self):
        self.shares = 0
        self.minimum = 0
        self.mode = ContextType.UNDEFINED
        pass

    def private_value(self, participant: int, value) -> int:
        """
        Return the private value for current context participant
        :param participant: current participant index in current share
        :param value: plaintext value to be encoded using secret sharing algorithms
        :return: encrypted value as result of secret sharing calculation
        """
        return value

    def reconstruct(self, value):
        """
        Return the decrypted value from given shares
        :param value: encrypted share value
        :return: decryptes value as result of secret sharing recovery
        """
        return value
