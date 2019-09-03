import torch
import torch.nn as nn
import torch.nn.functional as F
import torch.optim as optim
from torchvision import datasets, transforms
import syft as sy


class Arguments:
    def __init__(self):
        self.batch_size = 64
        self.test_batch_size = 200
        self.epochs = 10
        self.lr = 0.001  # learning rate
        self.log_interval = 100
        self.momentum = 0.5
        self.no_cuda = False
        self.seed = 1
        self.log_interval = 10
        self.save_model = True


class Net(nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.fc1 = nn.Linear(784, 500)
        self.fc2 = nn.Linear(500, 10)

    def forward(self, x):
        x = x.view(-1, 784)
        x = self.fc1(x)
        x = F.relu(x)
        x = self.fc2(x)
        return x


def train(args, model, train_loader, optimizer, epoch):
    model.train()
    for batch_idx, (data, target) in enumerate(train_loader):
        optimizer.zero_grad()
        output = model(data)
        output = F.log_softmax(output, dim=1)
        loss = F.nll_loss(output, target)
        loss.backward()
        optimizer.step()
        if batch_idx % args.log_interval == 0:
            print('Train Epoch: {} [{}/{} ({:.0f}%)]'.format(
                epoch, batch_idx * args.batch_size, len(train_loader) * args.batch_size, 100. * batch_idx / len(train_loader)))


def test(args, model, test_loader):
    model.eval()
    n_correct_priv = 0
    n_total = 0
    with torch.no_grad():
        for data, target in test_loader:
            output = model(data)
            pred = output.argmax(dim=1)
            n_correct_priv += pred.eq(target.view_as(pred)).sum()
            n_total += args.test_batch_size

            n_correct = n_correct_priv.copy().get().float_precision().long().item()

            print('Test set: Accuracy: {}/{} ({:.0f}%)'.format(
                n_correct, n_total,
                100. * n_correct / n_total))


def run_example():
    # Let’s start by importing the libraries and initializing the hook.
    # This is done to override PyTorch’s methods to execute commands on one
    # worker that are called on tensors controlled by the local worker. It
    # also allows us to move tensors between workers. Workers are explained below.
    hook = sy.TorchHook(torch)

    # load execution parameters
    args = Arguments()
    use_cuda = not args.no_cuda and torch.cuda.is_available()
    torch.manual_seed(args.seed)
    device = torch.device("cuda" if use_cuda else "cpu")
    kwargs = {'num_workers': 1, 'pin_memory': True} if use_cuda else {}

    # Virtual workers are entities present on our local machine.
    # They are used to model the behavior of actual workers.
    client = sy.VirtualWorker(hook, id="client")
    bob = sy.VirtualWorker(hook, id="bob")
    alice = sy.VirtualWorker(hook, id="alice")
    crypto_provider = sy.VirtualWorker(hook, id="crypto_provider")

    # download MNIST training dataset
    # Downloads MNIST dataset
    mnist_trainset = datasets.MNIST(
        root="../data",
        train=True,
        download=True,
        transform=transforms.Compose(
            [transforms.ToTensor(), transforms.Normalize((0.1307,), (0.3081,))]
        ),
    )
    # STRATEGY
    #
    # We can perform federated learning on client devices by following these steps:
    #
    #  send the model to the device,
    #  do normal training using the data present on the device,
    #  get back the smarter model.
    # However, if someone intercepts the smarter model while it is shared with the server,
    # he could perform reverse engineering and extract sensitive data about the dataset.
    # Differential privacy methods address this issue and protect the data.
    #
    # When the updates are sent back to the server, the server should not be able to
    # discriminate while aggregating the gradients. Let’s use a form of cryptography
    # called additive secret sharing.
    #
    # We want to encrypt these gradients (or model updates) before performing the aggregation
    # so that no one will be able to see the gradients. We can achieve this by additive secret sharing.

    # In our setting, we assume that the server has access to some data to first train its model. Here is the MNIST
    # training set.
    train_loader = torch.utils.data.DataLoader(
        datasets.MNIST('../data', train=True, download=True,
                       transform=transforms.Compose([
                           transforms.ToTensor(),
                           transforms.Normalize((0.1307,), (0.3081,))
                       ])),
        batch_size=args.batch_size, shuffle=True)
    # Second, the client has some data and would like to have predictions on it using the server's model. This
    # client encrypts its data by sharing it additively across two workers alice and bob.
    #
    # SMPC uses crypto protocols which require to work on integers. We leverage here the PySyft tensor
    # abstraction to convert PyTorch Float tensors into Fixed Precision Tensors using .fix_prec(). For
    # example 0.123 with precision 2 does a rounding at the 2nd decimal digit so the number stored is
    # the integer 12.
    test_loader = torch.utils.data.DataLoader(
        datasets.MNIST('../data', train=False,
                       transform=transforms.Compose([
                           transforms.ToTensor(),
                           transforms.Normalize((0.1307,), (0.3081,))
                       ])),
        batch_size=args.test_batch_size, shuffle=True)

    # Convert to integers and privately share the dataset
    private_test_loader = []
    for data, target in test_loader:
        shared_data = data.fix_prec().share(alice, bob, crypto_provider=crypto_provider)
        shared_target = target.fix_prec().share(alice, bob, crypto_provider=crypto_provider)
        private_test_loader.append((
            shared_data,
            shared_target
        ))
        print(shared_data)
        print(shared_target)
    # train the network
    model = Net()
    optimizer = torch.optim.Adam(model.parameters(), lr=args.lr)

    for epoch in range(1, args.epochs + 1):
        train(args, model, train_loader, optimizer, epoch)
    # Good, our model is now trained and ready to be provided as a service!
    print("model trained")
    # Now, as the server, we send the model to the workers holding the data.
    # Because the model is sensitive information (you've spent time optimizing it!),
    # you don't want to disclose its weights so you secret share the model just like
    # the client did with the test dataset earlier
    shared_model = model.fix_precision().share(alice, bob, crypto_provider=crypto_provider)
    print(alice)
    print(bob)
    # The following test function performs the encrypted evaluation. The model weights, the
    # data inputs, the prediction and the target used for scoring are all encrypted!
    # However as you can observe, the syntax is very similar to normal PyTorch testing! Nice!
    # The only thing we decrypt from the server side is the final score at the end of our
    # 200 items batches to verify predictions were on average good.
    test(args, model, private_test_loader)
    # Et voilà! Here you are, you have learned how to do end to end secure predictions:
    # the weights of the server's model have not leaked to the client and the server has
    # no information about the data input nor the classification output!
    #
    # Regarding performance, classifying one image takes less than 0.1 second, approximately
    # 33ms on my laptop (2,7 GHz Intel Core i7, 16GB RAM). However, this is using very fast
    # communication (all the workers are on my local machine). Performance will vary depending
    # on how fast different workers can talk to each other.


# example made from
# https://blog.openmined.org/encrypted-deep-learning-classification-with-pysyft/
if __name__ == '__main__':
    run_example()
