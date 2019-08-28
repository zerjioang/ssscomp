package common

import (
	"bytes"
	"math/big"
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 50000; i++ {
		if RandomPrime().Cmp(prime) >= 0 {
			t.Fatal("generated random numbers failed")
		}
	}
}

func TestBaseConversion(t *testing.T) {
	for i := 0; i < 10000; i++ {
		point := RandomPrime()
		if point.Cmp(FromBase64(ToBase64(point))) != 0 {
			t.Fatal("Fatal: Base conversion failed")
		}
	}
}

func TestToBase64(t *testing.T) {
	for i := 0; i < 10000; i++ {
		point := RandomPrime()
		if len(ToBase64(point)) != 44 {
			t.Fatal("Fatal: ToBase64 returned wrong length")
		}
	}
}

func TestSplitMerge(t *testing.T) {
	// Short, medium, and long tests
	tests := [][]byte{
		[]byte("N17FigASkL6p1EOgJhRaIquQLGvYV0"),
		[]byte("0y10VAfmyH7GLQY6QccCSLKJi8iFgpcSBTLyYOGbiYPqOpStAf1OYuzEBzZR"),
		[]byte("KjRHO1nHmIDidf6fKvsiXWcTqNYo2U9U8juO94EHXVqgearRISTQe0zAjkeUYYBvtcB8VWzZHYm6ktMlhOXXCfRFhbJzBUsXaHb5UDQAvs2GKy6yq0mnp8gCj98ksDlUultqygybYyHvjqR7D7EAWIKPKUVz4of8OzSjZlYg7YtCUMYhwQDryESiYabFID1PKBfKn5WSGgJBIsDw5g2HB2AqC1r3K8GboDN616Swo6qjvSFbseeETCYDB3ikS7uiK67ErIULNqVjf7IKoOaooEhQACmZ5HdWpr34tstg18rO"),
	}

	for i := range tests {
		result := MergeIntToByte(SplitByteToInt(tests[i]))
		if !bytes.Equal(result, tests[i]) {
			t.Fatal("Fatal: splitting and merging returned invalid data: ", result, tests[i])
		}
	}
}

func TestSplitMergeOdds(t *testing.T) {
	// Short, medium, and long tests
	tests := [][]byte{
		[]byte("a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00a"),
		[]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa哈囉世界"),
		[]byte("こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界こんにちは、世界"),
	}

	for i := range tests {
		result := MergeIntToByte(SplitByteToInt(tests[i]))
		if !bytes.Equal(result, tests[i]) {
			t.Fatal("Fatal: splitting and merging returned invalid data on test", i)
		}
	}
}

func TestModInverse(t *testing.T) {
	prime := GetDefaultPrimeNumber()
	for i := 0; i < 10000; i++ {
		point := big.NewInt(0).Set(RandomPrime())
		if point.Cmp(prime) >= 0 {
			t.Fatal("Error! Random point out of bounds exception")
		}
		inverse := big.NewInt(0).Set(ModInverse(point))
		if inverse.Cmp(prime) >= 0 {
			t.Fatal("Error! Inverse out of bounds exception")
		}
		value := big.NewInt(0).Set(point)
		value = value.Mul(value, inverse)
		value = value.Mod(value, prime)
		expected := big.NewInt(1)
		if value.Cmp(expected) != 0 {
			t.Fatalf("Fatal: ModInverse[%v] failed\nExpected: %v; Got: %v\nPoint %v\nInverse: %v\nPrime: %v", i, expected, value, point, inverse, prime)
		}
	}
}

func TestEvaluatePolynomial(t *testing.T) {
	values := [][][]*big.Int{
		{
			{big.NewInt(20), big.NewInt(21), big.NewInt(42)},
			{big.NewInt(0)},
		},
		{
			{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
			{big.NewInt(4)},
		},
		{
			{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)},
			{big.NewInt(10)},
		},
	}

	actual := []*big.Int{big.NewInt(20), big.NewInt(0), big.NewInt(54321)}

	for i := range values {
		result := EvaluatePolynomial(values[i][0], values[i][1][0])
		if result.Cmp(actual[i]) != 0 {
			t.Fatalf("Fatal: EvaluatePolynomial[%v] failed\nExpected: %v; Got: %v\n", i, actual[i], result)
		}
	}
}
