package sss

import (
	"errors"

	"github.com/zerjioang/ssscomp/lib/common"
)

var (
	errInvalidShares = errors.New("invalid number of shares provided")
)

func ShareAdd(a common.Shareable, b common.Shareable) (common.Shareable, error) {
	if a == nil || b == nil {
		return nil, errInvalidShares
	}
	return a.Add(b)
}

func ShareNegation(a common.Shareable) (common.Shareable, error) {
	if a == nil {
		return nil, errInvalidShares
	}
	return a.Neg()
}

func ShareSubstraction(a common.Shareable, b common.Shareable) (common.Shareable, error) {
	if a == nil || b == nil {
		return nil, errInvalidShares
	}
	return a.Sub(b)
}

func ShareDivision(a common.Shareable, b common.Shareable) (common.Shareable, error) {
	if a == nil || b == nil {
		return nil, errInvalidShares
	}
	return a.Div(b)
}

func SharePow(a common.Shareable, b int) (common.Shareable, error) {
	if a == nil {
		return nil, errInvalidShares
	}
	return a.Pow(b)
}

func ShareMul(a common.Shareable, b int) (common.Shareable, error) {
	if a == nil {
		return nil, errInvalidShares
	}
	return a.Mul(b)
}
