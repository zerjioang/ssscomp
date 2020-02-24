package main

import (
	"math/big"
	"testing"

	"github.com/zerjioang/ssscomp/lib/bigconst"
)

func BenchmarkCL15(b *testing.B) {
	b.Run("generate-prime-2048", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = bigconst.LargePrime(bits)
		}
	})
	b.Run("example", func(b *testing.B) {
		// number uid processing
		uid := new(big.Int).SetInt64(19382983298)
		xa := new(big.Int).SetInt64(3929333233)
		xb := new(big.Int).SetInt64(2389239238)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = generateIds(nil, nil, uid, xa, xb, bits)
		}
	})
	b.Run("prime-reuse", func(b *testing.B) {
		// number uid processing
		p, _ := bigconst.LargePrime(bits)
		q, _ := bigconst.LargePrime(bits)
		uid := new(big.Int).SetInt64(19382983298)
		xa := new(big.Int).SetInt64(3929333233)
		xb := new(big.Int).SetInt64(2389239238)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = generateIds(p, q, uid, xa, xb, bits)
		}
	})
}
