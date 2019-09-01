package simple

import (
	"errors"
	"fmt"
	"github.com/zerjioang/ssscomp/lib/common"
)

var (
	errMoreParticipantsRequired = errors.New("participant count must be bigger or equal than 2")
	errSchemaNotDefined         = errors.New("schema not defined")
)

/*
Additive Sharing

Let’s first assume that we have fixed a finite field to which all secrets
and shares belong, and in which all computation take place; this could for
instance be the integers modulo a prime number, i.e. { 0, 1, ..., Q-1 } for a prime Q.

An easy way to split a secret x from this field into say three shares
x1, x2, x3, is to simply pick x1 and x2 at random and let x3 = x - x1 - x2.
As argued below, this hides the secret as long as no one knows more than two shares,
yet if all three shares are known then x can be reconstructed by simply computing x1 + x2 + x3.
More generally, this scheme is known as additive sharing and works for any N number of shares
by picking T = N - 1 random values.

*/

/*
logically, we must have R <= N since otherwise reconstruction is never possible
and we must have T < R since otherwise privacy makes little sense.

SimpleAdditiveScheme is an n-out-of-n schema
(It’s clear that this scheme isn-out-of-n, since any n−1 shares are random,
so allnshares arerequired to recover the secret.)
*/
type SimpleAdditiveScheme struct {
	common.SecretSchema `json:"_,omitempty"`
	// N: the number of shares that each secret is split into
	N int `json:"n"`
	// R: the minimum number of shares needed to reconstruct the secret
	R int `json:"r"`
	// T: the maximum number of shares that may be seen without
	// learning nothing about the secret, also known as the privacy threshold
	T int `json:"t"`
	// K: the number of secrets shared together
	K int `json:"k"`
	// Q limit of the finite space
	Q int `json:"q"`
}

func NewSimpleAdditiveScheme(participants int) (common.SecretSchema, error) {
	ret := new(SimpleAdditiveScheme)
	if participants < 2 {
		return nil, errMoreParticipantsRequired
	}
	ret.N = participants
	ret.R = ret.N
	ret.K = 1
	ret.T = ret.R - ret.K
	// pick a random Q (the integers modulo a prime number)
	ret.Q = common.RandomInRange(0, 5000000) + 1 // common.RandomInt()
	return ret, nil
}

func (as *SimpleAdditiveScheme) String() string {
	return fmt.Sprintf("simple sharing scheme: N=%d, R=%d, T=%d, K=%d, Q=%d", as.N, as.R, as.T, as.K, as.Q)
}

// returns a random number inside following finite space: { 0, 1, ..., Q-1 } for a prime Q.
func (as *SimpleAdditiveScheme) Random() int {
	return common.RandomInRange(0, as.Q)
}

func (as *SimpleAdditiveScheme) Generate(secret int) (shares []common.Shareable) {
	result := make([]common.Shareable, as.MinShares())
	var sum int
	for i := 0; i < int(as.N-1); i++ {
		// simply pick values at random for x1, x2, x3, x4, ...x(n-1)
		cv := as.Random()
		result[i] = common.NewIntSharePtr(cv)
		sum += cv
	}
	// set final value to x(n-1) = x - x1 - x2 - x3 - x4
	last := (secret - sum) % as.Q
	result[as.N-1] = common.NewIntSharePtr(last)
	return result
}

// return number of shared
func (as *SimpleAdditiveScheme) Shares() int {
	return as.N
}

func (as *SimpleAdditiveScheme) MinShares() int {
	return as.R
}

// return number of participants required to decode the message M
func (as *SimpleAdditiveScheme) PrivacyThreshold() int {
	return as.T
}

// if all three shares are known then x
// can be reconstructed by simply computing x1 + x2 + x3 + x(n+1)
// That the secret remains hidden as long as at most T = N - 1 shareholders
// collaborate follows from the marginal distribution of the view of up to T shareholders
// being independent of the secret.
func (as *SimpleAdditiveScheme) Reconstruct(shares []common.Shareable) (common.Shareable, error) {
	if as != nil {
		reconstructed := shares[0].Copy()
		reconstructed.Reset()
		for i := 0; i < len(shares); i++ {
			reconstructed, _ = reconstructed.Add(shares[i])
		}
		return reconstructed.Mod(as.Q)
	}
	return nil, errSchemaNotDefined
}
