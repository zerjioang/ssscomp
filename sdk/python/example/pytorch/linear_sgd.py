# Import Numpy & PyTorch
import numpy as np
import pandas as pd
import torch
import matplotlib.pyplot as plt
from torch.autograd import Variable
from torch.nn import Linear, Module, MSELoss
from torch.optim import SGD


class LinearRegressionModel(torch.nn.Module):

    def __init__(self, input_size, output_size):
        super(LinearRegressionModel, self).__init__()
        self.linear = Linear(input_size, output_size)  # One in and one out

    def forward(self, x):
        # x = x.view(x.size(0), -1)
        y_pred = self.linear(x)
        return y_pred

    # our model


def generate_dataset(size=100):
    # equation: y = 3x + 4
    inputs = []
    outputs = []
    for ix in range(size):
        # its important to scale the data to a value between 0 and 1
        random_number = np.random.randint(100) / 100
        inputs.append([random_number])
        outputs.append([3 * random_number + 4])
    return inputs, outputs


def run_example():
    # load example data
    cars_frame = pd.read_csv('../data/cars.csv', usecols=['speed', 'dist'])

    use_random_dataset = False
    if use_random_dataset:
        inputs, labels = generate_dataset(100)
    else:
        # extract columns
        inputs = cars_frame["speed"].values.tolist()
        inputs = np.array(inputs, dtype=np.float32)
        inputs = inputs.reshape(-1, 1)
        labels = cars_frame["dist"].values.tolist()
        labels = np.array(labels, dtype=np.float32)
        labels = labels.reshape(-1, 1)
        # data visualization
        plt.scatter(inputs, labels)
        plt.xlabel("speed")
        plt.ylabel("distance")
        plt.title("card dataset")
        plt.show()

    # parameters of our module
    input_dim = 1  # takes variable 'x'
    output_dim = 1  # takes variable 'y'
    learning_rate = 0.01
    momentum = 0.01
    epochs = 100

    # create our model
    linear_model = LinearRegressionModel(input_dim, output_dim)

    # For GPU
    if torch.cuda.is_available():
        linear_model.cuda()

    criterion = MSELoss()
    optimizer = SGD(linear_model.parameters(), lr=learning_rate, momentum=momentum)

    for epoch in range(epochs):
        epoch += 1
        # Convert numpy array to torch Variable
        i = torch.from_numpy(inputs).requires_grad_()
        l = torch.from_numpy(labels)

        # Clear gradients w.r.t. parameters
        optimizer.zero_grad()

        # Forward to get output
        outputs = linear_model(i)

        # Calculate Loss
        loss = criterion(outputs, l)

        # Getting gradients w.r.t. parameters
        loss.backward()

        # Updating parameters
        optimizer.step()

        print('epoch {}, loss {}'.format(epoch, loss.item()))
        '''
    for epoch in range(epochs):
        epoch += 1
        # iterate over each input
        epoch_loss = 0
        for ix, x in enumerate(inputs):
            # Converting inputs and labels to Variable
            if torch.cuda.is_available():
                inputs_tensor = Variable(torch.from_numpy(inputs[ix]).cuda())
            else:
                # here x is the input. i.e. the input value of x
                # and y_train[ix] is the output. i.e. y = f(x) = 3x + 4
                inputs_tensor = Variable(torch.Tensor(inputs))

            # Forward pass: Compute predicted y by passing
            # x to the model
            y_pred = linear_model(inputs_tensor)

            # Compute and print loss
            loss = criterion(y_pred, Variable(torch.Tensor([labels[ix]])))

            # Clear gradient buffers because we don't want any gradient
            # from previous epoch to carry forward, dont want to cummulate gradients
            optimizer.zero_grad()
            # get gradients w.r.t to parameters
            loss.backward()
            # update parameters
            optimizer.step()
            print('epoch {}, loss {}'.format(epoch, loss.item()))
'''

    # execute a prediction on trained model
    # test and plot our model
    linear_model.eval()
    if use_random_dataset:
        with torch.no_grad():
            v = Variable(torch.Tensor([12]))
            predicted = linear_model(v)
            print(predicted)
    else:
        test_inputs = inputs
        with torch.no_grad():
            # we don't need gradients in the testing phase
            for ix, x in enumerate(test_inputs):
                if torch.cuda.is_available():
                    predicted = linear_model(Variable(torch.from_numpy(test_inputs[ix]).cuda())).cpu().data.numpy()
                else:
                    predicted = linear_model(Variable(torch.from_numpy(test_inputs[ix]))).data.numpy()
                print(predicted)
        plt.clf()
        plt.plot(inputs, labels, 'go', label='True data', alpha=0.5)
        plt.plot(inputs, predicted, '--', label='Predictions', alpha=0.5)
        plt.legend(loc='best')
        plt.show()


# based on https://www.geeksforgeeks.org/linear-regression-using-pytorch/
# more info: https://towardsdatascience.com/linear-regression-with-pytorch-eb6dedead817
# https://linuxhint.com/pytorch_linear_regression_tutorial/
# https://www.deeplearningwizard.com/deep_learning/practical_pytorch/pytorch_linear_regression/
# https://medium.com/@ally_20818/pytorch-101-linear-regression-with-pytorch-d2d22291c37d
# https://towardsdatascience.com/private-ai-federated-learning-with-pysyft-and-pytorch-954a9e4a4d4e
# https://github.com/OpenMined/PySyft/blob/master/examples/tutorials/Part%2004%20-%20Federated%20Learning%20via%20Trusted%20Aggregator.ipynb
if __name__ == '__main__':
    run_example()
