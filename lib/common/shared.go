package common

// Secret Sharing item interface
type Shareable interface {
	IntValue() int
	// operations over current finite field (aka sample value)
	Add(shareable Shareable) (Shareable, error)
	Sub(shareable Shareable) (Shareable, error)
	Mul(b int) (Shareable, error)
	Pow(exponent int) (Shareable, error)
	Div(shareable Shareable) (Shareable, error)
	Neg() (Shareable, error)
	Reset()
	Copy() Shareable
	Mod(q int) (Shareable, error)
}

// SecretSchema is the interface used by all different secret sharing schemas
type SecretSchema interface {
	Shares() int
	MinShares() int
	PrivacyThreshold() int
	// reconstruct schema shares
	Reconstruct(shares []Shareable) (Shareable, error)
	//generates schema shares
	Generate(secret int) (shares []Shareable)
}

type HomomorphicSchema interface {
	Generate() error
	Encrypt()
	Decrypt()
}
