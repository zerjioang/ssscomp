package simple

import (
	"errors"
	"github.com/zerjioang/ssscomp/lib/common"
)

func AdditiveAdd(a common.Shareable, b common.Shareable) (common.Shareable, error) {
	if a == nil || b == nil {
		return nil, errors.New("invalid number of shares provided")
	}
	return a.Add(b)
}

func AdditiveNegation(a common.Shareable) (common.Shareable, error) {
	if a == nil {
		return nil, errors.New("invalid number of shares provided")
	}
	return a.Neg()
}

func AdditiveSubstraction(a common.Shareable, b common.Shareable) (common.Shareable, error) {
	if a == nil || b == nil {
		return nil, errors.New("invalid number of shares provided")
	}
	return a.Sub(b)
}

func AdditiveDivision(a common.Shareable, b common.Shareable) (common.Shareable, error) {
	if a == nil || b == nil {
		return nil, errors.New("invalid number of shares provided")
	}
	return a.Div(b)
}

func AdditivePow(a common.Shareable, b int) (common.Shareable, error) {
	if a == nil {
		return nil, errors.New("invalid number of shares provided")
	}
	return a.Pow(b)
}

func AdditiveMul(a common.Shareable, b int) (common.Shareable, error) {
	if a == nil {
		return nil, errors.New("invalid number of shares provided")
	}
	return a.Mul(b)
}
