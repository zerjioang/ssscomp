import torch
import torchvision
from torch import nn
import torch.optim as optim
import torch.nn.functional as F
from torchvision import datasets, transforms
import syft as sy


class Model(nn.Module):
    def __init__(self):
        super(Model, self).__init__()
        self.fc1 = nn.Linear(784, 500)
        self.fc2 = nn.Linear(500, 10)

    def forward(self, x):
        x = x.view(-1, 784)
        x = self.fc1(x)
        x = F.relu(x)
        x = self.fc2(x)
        return F.log_softmax(x, dim=1)


def run_example():
    # Let’s start by importing the libraries and initializing the hook.
    # This is done to override PyTorch’s methods to execute commands on one
    # worker that are called on tensors controlled by the local worker. It
    # also allows us to move tensors between workers. Workers are explained below.
    hook = sy.TorchHook(torch)

    # define virtual workers for federated learning
    jake = sy.VirtualWorker(hook, id="jake")
    print("Jake has: " + str(jake._objects))
    john = sy.VirtualWorker(hook, id="john")
    print("John has: " + str(john._objects))
    secure_worker = sy.VirtualWorker(hook, id="secure_worker")

    # STRATEGY
    # We can perform federated learning on client devices by following these steps:
    #  send the model to the device,
    #  do normal training using the data present on the device,
    #  get back the smarter model.

    # In real-life applications, the data is present on client devices.
    # To replicate the scenario, we send data to the VirtualWorkers.
    transform = transforms.Compose([
        transforms.ToTensor(),
        transforms.Normalize((0.5,), (0.5,)),
    ])

    train_set = datasets.MNIST(
        "../data", train=True, download=True, transform=transform)
    test_set = datasets.MNIST(
        "../data", train=False, download=True, transform=transform)

    # Notice that we have created the training dataset differently.
    # The train_set.federate((jake, john)) creates a FederatedDataset wherein
    # the train_set is split among Jake and John (our two VirtualWorkers).
    # The FederatedDataset class is intended to be used like the PyTorch’s Dataset class.
    # Pass the created FederatedDataset to a federated data loader “FederatedDataLoader”
    # to iterate over it in a federated manner. The batches then come from different devices.
    federated_train_loader = sy.FederatedDataLoader(
        train_set.federate((jake, john)), batch_size=64, shuffle=True)

    test_loader = torch.utils.data.DataLoader(
        test_set, batch_size=64, shuffle=True)

    # define the model
    model = Model()
    optimizer = optim.SGD(model.parameters(), lr=0.01)
    # train the model
    # Since the data is present on the client device, we obtain its location through the location attribute.
    # The important additions to the code are the steps to get back the improved model and the value of the
    # loss from the client devices.
    for epoch in range(0, 5):
        model.train()
        for batch_idx, (data, target) in enumerate(federated_train_loader):
            # send the model to the client device where the data is present
            model.send(data.location)
            # training the model
            optimizer.zero_grad()
            output = model(data)
            loss = F.nll_loss(output, target)
            loss.backward()
            optimizer.step()
            # get back the improved model
            model.get()
            if batch_idx % 100 == 0:
                # get back the loss
                loss = loss.get()
                print('Epoch: {:2d} [{:5d}/{:5d} ({:3.0f}%)]\tLoss: {:.6f}'.format(
                    epoch + 1,
                    batch_idx * 64,
                    len(federated_train_loader) * 64,
                    100. * batch_idx / len(federated_train_loader),
                    loss.item()))
    # test the model
    model.eval()
    test_loss = 0
    correct = 0
    with torch.no_grad():
        for data, target in test_loader:
            output = model(data)
            test_loss += F.nll_loss(
                output, target, reduction='sum').item()
            # get the index of the max log-probability
            pred = output.argmax(1, keepdim=True)
            correct += pred.eq(target.view_as(pred)).sum().item()

    test_loss /= len(test_loader.dataset)

    print('\nTest set: Average loss: {:.4f}, Accuracy: {}/{} ({:.0f}%)\n'.format(
        test_loss,
        correct,
        len(test_loader.dataset),
        100. * correct / len(test_loader.dataset)))
    # That’s it. We have trained a model using the federated learning approach.
    # When compared to traditional training, it takes more time to train a model using the federated approach.

    # PROTECTING THE MODEL
    #
    # Training the model on the client device protected the user’s privacy.
    # But, what about the model’s privacy? Downloading the model can threaten
    # the organization’s intellectual property!
    # Secure Multi-Party Computation, which consists of secret additive sharing,
    # provides us with a way to perform model training without disclosing the model.
    # To protect the weights of the model, we secret share the model among the client devices.
    shared_model = model.fix_precision().share(jake, john, crypto_provider=secure_worker)
    print(shared_model)

    # now the model, the inputs, model outputs, weights, etc.
    # will be encrypted as well. Working on encrypted inputs will yield encrypted output.


# example made from
# https://towardsdatascience.com/introduction-to-federated-learning-and-privacy-preservation-75644686b559
if __name__ == '__main__':
    run_example()
