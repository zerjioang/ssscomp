package main

import (
	"fmt"
	"github.com/sajari/regression"
	"github.com/zerjioang/ssscomp/lib/helper"
	"github.com/zerjioang/ssscomp/lib/toy"
)

func main() {
	bobSchema := toy.NewIntegerToyHomoScheme()
	_ = bobSchema.Generate()

	// Bob generates a LM
	r := new(regression.Regression)
	r.SetObserved("dist")
	r.SetVar(0, "speed")
	_ = helper.LoadCsv("./sdk/python/example/data/cars.csv", 3, true, func(obs float64, vars []float64) {
		r.Train(
			regression.DataPoint(
				bobSchema.EncryptF64F(obs),
				bobSchema.EncryptF64FArray(vars)))
	})
	_ = r.Run()
	/*
		N = 50
		Variance observed = 650.7796000000006
		Variance Predicted = 423.7091789780147
		R2 = 0.6510793807581158
	*/
	fmt.Printf("Bob Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Bob Regression:\n%s\n", r)
	// safety distance prediction
	// Predicted = -50924.20 + speed*3.93
	// Linear regression formula for encrypted (additive sharded) data: -50924.20 + speed*3.93
	// Bob makes a prediction
	prediction, err := r.Predict([]float64{bobSchema.EncryptF64F(20)})
	if err == nil {
		fmt.Println("Bob Additive Shard Encrypted Prediction: ", prediction)
		plain := bobSchema.Decrypt(int(prediction))
		fmt.Println("Bob Decrypted Prediction: ", plain)

	}
	fmt.Println("In order to Alice to Predict over encrypted data, Bob sends its N value to alice for data pre-processing")
	// for 20 (speed), a 61 distance is predicted
	// Alice wants to use encrypted model with its own data
	encryptedAliceValue := bobSchema.N()+20
	// alice prediction over encrypted data
	alicePrediction, err := r.Predict([]float64{float64(encryptedAliceValue)})
	if err == nil {
		fmt.Println("Alice Additive Shard Encrypted Prediction: ", alicePrediction)
		plain := bobSchema.Decrypt(int(alicePrediction))
		fmt.Println("Alice Decrypted Prediction: ", plain)

	}
}
