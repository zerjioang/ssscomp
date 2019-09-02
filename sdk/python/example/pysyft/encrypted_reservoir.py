import torch
import syft as sy
import numpy as np
import matplotlib.pyplot as plt

# demonstrate secure multiparty computation
display_figs = True
Q = 501337


def encrypt(x):
    share_a = np.random.randint(0, Q)
    share_b = np.random.randint(0, Q)
    share_c = (x - share_a - share_b) % Q
    return share_a, share_b, share_c


def decrypt(*shares):
    return sum(shares) % Q


def add(x, y):
    z = []
    assert (len(x) == len(y)), 'each variable must consist of the same number of shares!'
    for ii in range(len(x)):
        z.append((x[ii] + y[ii]) % Q)
    return z


def product(x, w):
    # w is a plaintext value,
    z = []
    for ii in range(len(x)):
        z.append((x[ii] * w) % Q)
    return z


if __name__ == '__main__':
    hook = sy.TorchHook(torch)
    # encrypt variables
    print('Encrypting variables 2 and 5 as var1 and var2')
    var1 = encrypt(2)
    var2 = encrypt(5)
    # get sum
    print('Performing encrypted addition...')
    my_sum = add(var1, var2)
    # print results
    print('Multiparty result from add(var1,var2):\n\t\t\t\t', my_sum)
    print('Decrypted result:\n\t\t\t\t', decrypt(*my_sum))
    print('Decrypted partial result:\n\t\t\t\t', decrypt(*my_sum[0:2]))
    bob = sy.VirtualWorker(hook, id='bob')
    alice = sy.VirtualWorker(hook, id='alice')
    # Create our dataset: an XOR truth table
    x = np.array([[0., 0], [0, 1], [1, 0], [1, 1]], 'float')
    y = np.array([[0], [1], [1], [0]], 'float')
    # use a reservoir transformation to achieve non-linearity in the model
    res_size = 256
    my_transform = np.random.randn(2, res_size)
    x_2 = np.matmul(x, my_transform)
    # apply relu non_linearity
    x_2[x_2 < 0] = 0.
    # convert data and targets to Syft tensors
    data = sy.FloatTensor(x_2)  # [[0,0],[0,1],[1,0],[    1,1]])
    target = sy.FloatTensor(y)  # [[0],[1],[1],[0]])
    # init model (just a matrix for linear regression)
    model = sy.zeros(res_size, 1)
    # encrypt and share the data, targets, and model
    data = data.fix_precision().share(alice, bob)
    target = target.fix_precision().share(alice, bob)
    model = model.fix_precision().share(alice, bob)
    # train the model
    learning_rate = 1e-3
    J = []
    print('\nBegin training encrypted reservoir...')
    for i in range(250):
        pred = data.mm(model)
        grad = pred - target
        update = data.transpose(0, 1).mm(grad)
        model = model - update * learning_rate
        loss = grad.get().decode().abs().sum()
        J.append(loss)
        if i % 50 == 0:
            print('loss at step %i: %.3e' % (i, loss))
    got_pred = pred.get()
    got_target = target.get()

    if display_figs:
        # display training results
        plt.figure(figsize=(10, 6))
        plt.plot(J, 'g', lw=4)
        plt.xlabel('step', fontsize=19)
        plt.ylabel('loss', fontsize=19)
        plt.title('Learning XOR While Encrypted', fontsize=20)
        plt.show()
    # print decrypted predictions and targets (decision boundary of predictions at 0.5)
    print('predictions: \n', [got_pred.decode() > 0.5])
    print('targets: \n', got_target.decode())
