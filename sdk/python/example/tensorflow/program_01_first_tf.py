import tensorflow as tf


def program_01():
    """
    Tensorflow example program that prepares a job to sum two values (3, 2)
    :return:
    """
    a = tf.constant(3)
    b = tf.constant(2)

    c = a + b

    sess = tf.Session()

    print(sess.run(c))


if __name__ == '__main__':
    program_01()
