package simple

import (
	"errors"
	"github.com/zerjioang/s3go/lib/common"
	"math"
)

func AdditiveAdd(a common.SecretNumber, b common.SecretNumber) (common.SecretNumber, error) {
	sum := NewSimpleAdditiveScheme(a.Shares())
	if a.Shares() != b.Shares() {
		return sum, errors.New("invalid number of participants in given samples")
	}
	sum.shares = make([]int, a.MinShares())
	for i := 0; i < a.MinShares(); i++ {
		sum.shares[i] = a.Split(i) + b.Split(i)
	}
	return sum, nil
}

func AdditiveNegation(a common.SecretNumber) common.SecretNumber {
	sum := NewSimpleAdditiveScheme(a.Shares())
	sum.shares = make([]int, a.MinShares())
	for i := 0; i < a.MinShares(); i++ {
		sum.shares[i] = -a.Split(i)
	}
	return sum
}

func AdditiveSubstraction(a common.SecretNumber, b common.SecretNumber) (common.SecretNumber, error) {
	sum := NewSimpleAdditiveScheme(a.Shares())
	if a.Shares() != b.Shares() {
		return sum, errors.New("invalid number of participants in given samples")
	}
	sum.shares = make([]int, a.MinShares())
	for i := 0; i < a.MinShares(); i++ {
		sum.shares[i] = a.Split(i) - b.Split(i)
	}
	return sum, nil
}

func AdditiveDivision(a common.SecretNumber, b common.SecretNumber) (common.SecretNumber, error) {
	sum := NewSimpleAdditiveScheme(a.Shares())
	if a.Shares() != b.Shares() {
		return sum, errors.New("invalid number of participants in given samples")
	}
	sum.shares = make([]int, a.MinShares())
	for i := 0; i < a.MinShares(); i++ {
		d := float64(a.Split(i)) / float64(b.Split(i))
		sum.shares[i] = int(math.Round(d * 10000000000000000))
	}
	return sum, nil
}
func AdditiveComparison(a common.SecretNumber, b common.SecretNumber) (common.SecretNumber, error) {
	result, err := AdditiveSubstraction(a, b)
	if err != nil {
		return nil, err
	}
	return result, nil
}
