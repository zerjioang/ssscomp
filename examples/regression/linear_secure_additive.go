package main

import (
	"fmt"

	"github.com/sajari/regression"
	"github.com/zerjioang/ssscomp/lib/common"
	"github.com/zerjioang/ssscomp/lib/helper"
	"github.com/zerjioang/ssscomp/lib/simple"
)

func main() {
	bobSchema, _ := simple.NewSimpleAdditiveSchemePtr(3)

	// Bob generates a LM for each shares
	// todo fix encryptio n process in proper shares
	var sharedModels [3]*regression.Regression
	for i := range sharedModels {
		r := new(regression.Regression)
		r.SetObserved("dist")
		r.SetVar(0, "speed")
		_ = helper.LoadCsv("./sdk/python/example/data/cars.csv", 3, true, func(obs float64, vars []float64) {
			r.Train(
				regression.DataPoint(
					bobSchema.EncryptF64F(obs, i),
					bobSchema.EncryptF64FArray(vars, i)))
		})
		_ = r.Run()
		sharedModels[i] = r
		fmt.Printf("Share [%d] Regression formula:\n%v\n", i, r.Formula)
	}

	// predict
	var results []common.Shareable
	results = make([]common.Shareable, 3)
	for i := range sharedModels {
		r := sharedModels[i]
		// safety distance prediction
		// Predicted = -50924.20 + speed*3.93
		// Linear regression formula for encrypted (additive sharded) data: -50924.20 + speed*3.93
		// Bob makes a prediction
		prediction, err := r.Predict([]float64{bobSchema.EncryptF64F(20, i)})
		if err == nil {
			fmt.Printf("Share [%d] Encrypted Prediction: %v\n", i, prediction)
			results[i] = common.NewIntSharePtr(int(prediction))
		}
	}
	p, _ := bobSchema.Reconstruct(results)
	fmt.Printf("Decrypted Prediction: %d\n", p.IntValue())
}
