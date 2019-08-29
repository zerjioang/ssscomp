import tensorflow as tf
from lib import SSSComp


def program_01_sscomp():
    """
    Tensorflow example program that prepares a job to sum two values (3, 2)
    using secure computation algorithms
    :return:
    """
    # create a Secure Multiparty Computation context for 3 participants
    context = SSSComp().new_smpc_additive(3)

    # request a private version of the value to be used for computation
    a = tf.constant(context.private_value(3))
    b = tf.constant(context.private_value(2))

    c = a + b

    sess = tf.Session()

    result = sess.run(c)
    context.reconstruct(result)


if __name__ == '__main__':
    program_01_sscomp()
