package math

import (
	"testing"
)

var result int64

// BenchmarkOptimizations/horner-4         	200000000	         8.41 ns/op	 118.88 MB/s	       0 B/op	       0 allocs/op
func BenchmarkOptimizations(b *testing.B) {
	b.Run("horner", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		var r int64
		for n := 0; n < b.N; n++ {
			r = Horner(3, []int64{-19, 7, -4, 6})
		}
		result = r
	})
}
