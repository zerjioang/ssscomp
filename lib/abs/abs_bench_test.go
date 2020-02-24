package abs_test

import (
	"testing"

	"github.com/zerjioang/ssscomp/lib/abs"
)

// sink is used to prevent the compiler from dropping function calls where the
// returned value is not used within benchmarks.
var sink int64

// BenchmarkRand measures the overhead incurred by the RNG for other benchmarks.
func BenchmarkRand(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		sink += Rand()
	}
}

func BenchmarkWithBranch(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		sink += WithBranch(Rand())
	}
}

func BenchmarkWithStdLib(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		sink += WithStdLib(Rand())
	}
}

func BenchmarkWithTwosComplement(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		sink += WithTwosComplement(Rand())
	}
}

func BenchmarkWithASM(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		sink += abs.WithASM(Rand())
	}
}
