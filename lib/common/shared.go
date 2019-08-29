package common

// SecretNumber is the interface used by all different secret sharing schemas
type SecretNumber interface {
	Reconstruct(shares []Share) int
	Json(split int) ([]byte, error)
	Shares() int
	MinShares() int
	PrivacyThreshold() int
	Split(idx int) int
	Next() (Share, error)
}
