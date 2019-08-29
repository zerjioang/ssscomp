package simple

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zerjioang/s3go/lib/common"
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
	common.SecretNumber `json:"_,omitempty"`
	// N: the number of shares that each secret is split into
	N int `json:"n"`
	// R: the minimum number of shares needed to reconstruct the secret
	R int `json:"r"`
	// T: the maximum number of shares that may be seen without
	// learning nothing about the secret, also known as the privacy threshold
	T int `json:"t"`
	// K: the number of secrets shared together
	K            int `json:"k"`
	shares       []int
	CurrentShare int `json:"share"`
}

func NewSimpleAdditiveScheme(participants int) *SimpleAdditiveScheme {
	ret := new(SimpleAdditiveScheme)
	ret.N = participants
	ret.R = ret.N
	ret.K = 1
	ret.T = ret.R - ret.K
	ret.shares = make([]int, participants)
	return ret
}

// if all three shares are known then x
// can be reconstructed by simply computing x1 + x2 + x3 + x(n+1)
// That the secret remains hidden as long as at most T = N - 1 shareholders
// collaborate follows from the marginal distribution of the view of up to T shareholders
// being independent of the secret.
func (as *SimpleAdditiveScheme) Reconstruct(shares []int) int {
	if as != nil {
		var reconstructed int
		for i := 0; i < len(shares); i++ {
			reconstructed += shares[i]
		}
		fmt.Println(reconstructed)
		return reconstructed
	}
	return 0
}

func (as *SimpleAdditiveScheme) String() string {
	return fmt.Sprintf("simple sharing scheme: N=%d, R=%d, T=%d, K=%d [%+v]", as.N, as.R, as.T, as.K, as.shares)
}

// converts current scheme to json only revealing specified share
func (as *SimpleAdditiveScheme) Json(split int) ([]byte, error) {
	if split >= 0 && split < len(as.shares) {
		as.CurrentShare = as.shares[split]
	}
	return json.Marshal(as)
}

// return number of shared
func (as *SimpleAdditiveScheme) Shares() int {
	return as.N
}

func (as *SimpleAdditiveScheme) MinShares() int {
	return as.R
}

// return number of participants
func (as *SimpleAdditiveScheme) PrivacyThreshold() int {
	return as.T
}
func (as *SimpleAdditiveScheme) Split(idx int) int {
	return as.shares[idx]
}

// returns specified share for given participant
func (as *SimpleAdditiveScheme) Next() (common.Share, error) {
	if as.CurrentShare >= 0 && as.CurrentShare < as.MinShares() {
		cs := common.NewShare(as.shares[as.CurrentShare])
		as.CurrentShare++
		return cs, nil
	}
	return common.Share{}, errors.New("no more shares available")
}

func (as *SimpleAdditiveScheme) Explain(seenShares []int, guess int) {
	// compute the unseen share that justifies the seen shares and the guess
	//simulated_unseen_share := (guess - sum(seen_shares)) % q
}

// secret: the secret to be splitted among all N
// N: number of N that will receive the secret
func SimpleAdditiveSecret(secret int, participants int) (common.SecretNumber, error) {
	ret := NewSimpleAdditiveScheme(participants)
	var sum int
	for i := 0; i < int(participants-1); i++ {
		// simply pick values at random from x1, x2, x3, x4, ...x(n-1)
		ret.shares[i] = common.RandomInRange(2, 200) //common.RandomInt()
		sum += ret.shares[i]
	}
	// set final value to x(n-1) = x - x1 - x2 - x3 - x4
	ret.shares[participants-1] = secret - sum
	return ret, nil
}
