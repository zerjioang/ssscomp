data("cars")
head(cars)
names(cars)
plot(dist~speed, data = cars)
mean(cars$dist)
mean(cars$dist, na.rm = T)

# Create Training and Test data -
set.seed(100)  # setting seed to reproduce results of random sampling
trainingRowIndex <- sample(1:nrow(cars), 0.8*nrow(cars))  # row indices for training data
trainingData <- cars[trainingRowIndex, ]  # model training data
testData  <- cars[-trainingRowIndex, ]   # test data

model1 = lm(dist~speed, data = trainingData)

print(model1)
plot(model1)
termplot(model1)
summary(model1)
write.csv(cars, file = "cars.csv")

distPred <- predict(model1, testData)  # predict distance

# calculate accuracy
actuals_preds <- data.frame(cbind(actuals=testData$dist, predicteds=distPred))  # make actuals_predicteds dataframe.
correlation_accuracy <- cor(actuals_preds)  # 82.7%
head(actuals_preds)

# Min-Max Accuracy Calculation
min_max_accuracy <- mean(apply(actuals_preds, 1, min) / apply(actuals_preds, 1, max))
# => 38.00%, min_max accuracy

# MAPE Calculation: MeanAbsolutePercentageError
mape <- mean(abs((actuals_preds$predicteds - actuals_preds$actuals))/actuals_preds$actuals)
# => 69.95%, mean absolute percentage deviation
