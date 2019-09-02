package main

import (
	"fmt"
	"github.com/sajari/regression"
	"github.com/zerjioang/ssscomp/lib/toy"
	"math"
)

func main() {
	r := new(regression.Regression)
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")

	schema := toy.NewIntegerToyHomoScheme()
	_ = schema.Generate()
	r.Train(
		regression.DataPoint(schema.EncryptF64(112), []float64{
			schema.EncryptF64(5870000),
			schema.EncryptF64(165),
			schema.EncryptF64(62),
		}),
		regression.DataPoint(schema.EncryptF64(134), []float64{
			schema.EncryptF64(6430000),
			schema.EncryptF64(205),
			schema.EncryptF64(64),
		}),
		regression.DataPoint(schema.EncryptF64(407), []float64{
			schema.EncryptF64(6350000),
			schema.EncryptF64(263),
			schema.EncryptF64(93),
		}),
		regression.DataPoint(schema.EncryptF64(53), []float64{
			schema.EncryptF64(6920000),
			schema.EncryptF64(165),
			schema.EncryptF64(53),
		}),
		regression.DataPoint(schema.EncryptF64(248), []float64{
			schema.EncryptF64(12480000),
			schema.EncryptF64(192),
			schema.EncryptF64(73),
		}),
		regression.DataPoint(schema.EncryptF64(127), []float64{
			schema.EncryptF64(6430000),
			schema.EncryptF64(165),
			schema.EncryptF64(59),
		}),
		regression.DataPoint(schema.EncryptF64(209), []float64{
			schema.EncryptF64(19640000),
			schema.EncryptF64(202),
			schema.EncryptF64(64),
		}),
		regression.DataPoint(schema.EncryptF64(357), []float64{
			schema.EncryptF64(15310000),
			schema.EncryptF64(213),
			schema.EncryptF64(76),
		}),
		regression.DataPoint(schema.EncryptF64(87), []float64{
			schema.EncryptF64(7130000),
			schema.EncryptF64(172),
			schema.EncryptF64(49),
		}),
		regression.DataPoint(schema.EncryptF64(96), []float64{
			schema.EncryptF64(7490000),
			schema.EncryptF64(143),
			schema.EncryptF64(64),
		}),
		regression.DataPoint(schema.EncryptF64(145), []float64{
			schema.EncryptF64(78950000),
			schema.EncryptF64(181),
			schema.EncryptF64(6),
		}),
		regression.DataPoint(schema.EncryptF64(269), []float64{
			schema.EncryptF64(7620000),
			schema.EncryptF64(231),
			schema.EncryptF64(74),
		}),
		regression.DataPoint(schema.EncryptF64(157), []float64{
			schema.EncryptF64(27930000),
			schema.EncryptF64(191),
			schema.EncryptF64(58),
		}),
		regression.DataPoint(schema.EncryptF64(362), []float64{
			schema.EncryptF64(7410000),
			schema.EncryptF64(247),
			schema.EncryptF64(86),
		}),
		regression.DataPoint(schema.EncryptF64(181), []float64{
			schema.EncryptF64(6250000),
			schema.EncryptF64(186),
			schema.EncryptF64(65),
		}),
		regression.DataPoint(schema.EncryptF64(289), []float64{
			schema.EncryptF64(8540000),
			schema.EncryptF64(249),
			schema.EncryptF64(83),
		}),
		regression.DataPoint(schema.EncryptF64(149), []float64{
			schema.EncryptF64(7160000),
			schema.EncryptF64(179),
			schema.EncryptF64(67),
		}),
		regression.DataPoint(schema.EncryptF64(258), []float64{
			schema.EncryptF64(9210000),
			schema.EncryptF64(224),
			schema.EncryptF64(86),
		}),
		regression.DataPoint(schema.EncryptF64(217), []float64{
			schema.EncryptF64(5950000),
			schema.EncryptF64(202),
			schema.EncryptF64(84),
		}),
		regression.DataPoint(schema.EncryptF64(257), []float64{
			schema.EncryptF64(33530000),
			schema.EncryptF64(169),
			schema.EncryptF64(67),
		}),
	)
	_ = r.Run()
	/*
	N = 20
	Variance observed = 9276.009999999997
	Variance Predicted = 8153.820887259359
	R2 = 0.8790224339192564
	*/
	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)
	// Predicted = -20030563.66 + Inhabitants*0.00 + Percent with incomes below $5000*1.26 + Percent unemployed*4.66
	encryptedPrediction, err := r.Predict([]float64{
		schema.EncryptF64(587000),
		schema.EncryptF64(165),
		schema.EncryptF64(62),
	})
	if err == nil {
		fmt.Println("Prediction over encrypted data: ", encryptedPrediction)
		// Prediction using encrypted data: 804264.9323967863
		plainVal := schema.Decrypt(int(math.Round(encryptedPrediction)))
		fmt.Println("Decrypted prediction value: ", plainVal)
	}

	plainPrediction, err := r.Predict([]float64{5870000, 165, 62})
	if err == nil {
		fmt.Println("Plain Prediction: ", plainPrediction)
		r := schema.Decrypt(136904)
		fmt.Println("Decrypted prediction value: ", r)
	}
}
