package common

import "testing"

// BenchmarkIsPrime/is-prime-true-12         	  200000	      6415 ns/op	   0.16 MB/s	       0 B/op	       0 allocs/op
// BenchmarkIsPrime/is-prime-false-12        	50000000	        30.5 ns/op	  32.76 MB/s	       0 B/op	       0 allocs/op      0 allocs/op
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

// BenchmarkIsPrimeBig/is-prime-true-12         	    2000	    602388 ns/op	   0.00 MB/s	   33056 B/op	     468 allocs/op
// BenchmarkIsPrimeBig/is-prime-false-12        	300000000	         5.80 ns/op	 172.42 MB/s	       0 B/op	       0 allocs/op	       3 allocs/op
func BenchmarkIsPrimeBig(b *testing.B) {
	b.Run("is-prime-true", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsPrimeBig("170141183460469231731687303715884105727")
		}
	})
	b.Run("is-prime-false", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsPrimeBig("1701411834604692317316873037158841057270")
		}
	})
}
