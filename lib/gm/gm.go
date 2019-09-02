package gm

import (
	"github.com/zerjioang/ssscomp/lib/math"
	"math/big"
)

type GM struct {
	p *big.Int
	q *big.Int
	n *big.Int
}

func NewGMSchema() *GM {
	schema := new(GM)
	return schema
}

func (schema *GM) Generate() error {
	var pErr, qErr error
	schema.p, pErr = math.Prime(10)
	if pErr != nil {
		return pErr
	}
	schema.q, qErr = math.Prime(10)
	if qErr != nil {
		return qErr
	}
	// compute n = p * q
	schema.n = big.NewInt(0)
	schema.n.Mul(schema.p, schema.q)
}
