package common

import "testing"

// BenchmarkIsPrime/is-prime-true-12         	    2000	    722837 ns/op	   0.00 MB/s	       0 B/op	       0 allocs/op
// BenchmarkIsPrime/is-prime-false-12        	50000000	        29.5 ns/op	  33.92 MB/s	       0 B/op	       0 allocs/op
func BenchmarkIsPrime(b *testing.B) {
	b.Run("is-prime-true", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsPrime(59239)
		}
	})
	b.Run("is-prime-false", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsPrime(592390)
		}
	})
}
