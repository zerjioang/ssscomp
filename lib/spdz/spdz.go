package spdz

// implemented from references:
// * https://github.com/mortendahl/privateml/blob/master/spdz/Basic%20SPDZ.ipynb
var (
	// These define the fields in which we're doing all of our computations.
	// small and large field
	// < 64 bits
	Q = 6497992661811505123
	// < 270 bits
	P = 1802216888453791673313287943102424579859887305661122324585863735744776691801009887
)

// Public elements
// These will represent any value that we wish to use in our
// computation that is not required to be kept private -- in fact,
// we'll often assume that both parties know all public values.
// The reason for having these (as opposed to just using floats etc.)
// is that all values must be brought into our finite field before
// we can use them; this class takes care of doing that.
