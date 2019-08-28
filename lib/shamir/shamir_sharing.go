package shamir

import (
	"errors"
	"github.com/zerjioang/s3go/lib/common"
)

/*
In Shamir’s scheme, instead of picking random field elements that sum up to the secret x
as we did above, to share x we sample a random polynomial f with the condition that f(0) = x
and evaluate this polynomial at N non-zero points to obtain the shares as f(1), f(2), …, f(N).
*/
type ShamirSharingScheme struct {
	// N: the number of shares that each secret is split into
	N int
	// R: the minimum number of shares needed to reconstruct the secret
	R      int
	secret int
	shares []int
}

// tries to reconstruct given secret from a set of points
func (as *ShamirSharingScheme) Reconstruct(points []int) (error, int) {
	if len(points) != as.R {
		return errors.New("invalid number of points provided"), -1
	}
	//lgi := lagrange.New()
	return nil, 0
}

// split secret using shamir scheme into requested shares given that only 'minimumSplits' shares are needed to reconstruct
func (as *ShamirSharingScheme) Generate() {
	// 1 generate k - 1 random numbers
	k := as.R
	rset := common.RandomSet(k - 1)
	// We construct as many points as participants using the polynomial function
	shares := as.buildSharedFromPolynomial(rset)
	as.shares = shares
}

func (as *ShamirSharingScheme) buildSharedFromPolynomial(coeficients []int) []int {
	var result []int
	result = make([]int, as.N)
	// build our polynomial function to generate secret shares (points)
	// f(x) = secret + a1 * x + a2 * x + an * x
	for i := 0; i < as.N; i++ {
		result[i] = as.secret + common.Polynomial(i+1, coeficients)
	}
	return result
}

// it will print generated points using shamir secret schema
func (as *ShamirSharingScheme) Print() {
	for i := 0; i < len(as.shares); i++ {
		println("share (", i+1, ",f(", i+1, "))", as.shares[i])
	}
}

// The essential idea of Adi Shamir's threshold scheme is that
// 2 points are sufficient to define a line,
// 3 points are sufficient to define a parabola,
// 4 points to define a cubic curve and so forth.
// That is, it takes k points to define a polynomial of degree k − 1
func NewShamirSharingScheme(secret int, participants int, min int) ShamirSharingScheme {
	var scheme ShamirSharingScheme
	scheme.secret = secret
	// set participants involved in this scheme
	scheme.N = participants
	// R: the minimum number of shares needed to reconstruct the secret
	scheme.R = min
	return scheme
}
