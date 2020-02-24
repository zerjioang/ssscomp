package main

import (
	"fmt"

	"github.com/sajari/regression"
	"github.com/zerjioang/ssscomp/lib/helper"
)

func main() {
	r := new(regression.Regression)
	r.SetObserved("dist")
	r.SetVar(0, "speed")
	_ = helper.LoadCsv("./sdk/python/example/data/cars.csv", 3, true, func(obs float64, vars []float64) {
		r.Train(regression.DataPoint(obs, vars))
	})
	_ = r.Run()
	/*
		N = 50
		Variance observed = 650.7796000000001
		Variance Predicted = 423.7091789781027
		R2 = 0.6510793807582516
	*/
	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)
	// safety distance prediction
	// Linear regression formula: -17.58 + speed*3.93
	prediction, err := r.Predict([]float64{20})
	if err == nil {
		fmt.Println("Prediction: ", prediction)
	}
	// for 20 (speed), a 61 distance is predicted
}
