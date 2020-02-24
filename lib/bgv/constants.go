package bgv

const (
	// definition of pi value
	Pi = 3.14159265358979323846
	// other BGV constants
	BgvSigma = 2400 //8
	BgvMu    = 60
	BgvVar   = 26
)

type Gen struct {
}

type BGV struct {
	n, d, L, N                       int
	q, f, s, t, A, T, tensorProductS Gen
}
