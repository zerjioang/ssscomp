import torch
import torch.nn as nn
import torch.nn.functional as F
import torch.optim as optim
from torchvision import datasets, transforms


class Arguments:
    def __init__(self):
        self.batch_size = 64
        self.test_batch_size = 200
        self.epochs = 10
        self.lr = 0.001  # learning rate
        self.log_interval = 100


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
    n_correct = 0
    n_total = 0
    with torch.no_grad():
        for data, target in test_loader:
            output = model(data)
            pred = output.argmax(dim=1)
            n_correct += pred.eq(target.view_as(pred)).sum()
            n_total += args.test_batch_size

            print('Test set: Accuracy: {}/{} ({:.0f}%)'.format(
                n_correct, n_total,
                100. * n_correct / n_total))


def run_example():
    args = Arguments()
    mnist_trainset = datasets.MNIST(
        root="../data",
        train=True,
        download=True,
        transform=transforms.Compose(
            [transforms.ToTensor(), transforms.Normalize((0.1307,), (0.3081,))]
        ),
    )
    train_loader = torch.utils.data.DataLoader(
        datasets.MNIST('../data', train=True, download=True,
                       transform=transforms.Compose([
                           transforms.ToTensor(),
                           transforms.Normalize((0.1307,), (0.3081,))
                       ])),
        batch_size=args.batch_size, shuffle=True)
    test_loader = torch.utils.data.DataLoader(
        datasets.MNIST('../data', train=False,
                       transform=transforms.Compose([
                           transforms.ToTensor(),
                           transforms.Normalize((0.1307,), (0.3081,))
                       ])),
        batch_size=args.test_batch_size, shuffle=True)
    model = Net()
    optimizer = torch.optim.Adam(model.parameters(), lr=args.lr)

    # load or train
    if True:
        # Model class must be defined somewhere
        model = torch.load("../data/MNIST/models/plaintext_cnn.pt")
        model.eval()
        test(args, model, test_loader)
    else:
        for epoch in range(1, args.epochs + 1):
            train(args, model, train_loader, optimizer, epoch)
        print("model trained")
        test(args, model, test_loader)
        # save our plaintext model
        torch.save(model.state_dict(), "../data/MNIST/models/plaintext_cnn.pt")


# example made from
# https://blog.openmined.org/encrypted-deep-learning-classification-with-pysyft/
if __name__ == '__main__':
    run_example()
