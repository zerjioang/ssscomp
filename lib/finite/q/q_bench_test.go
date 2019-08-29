package q

import (
	"math/big"
	"testing"
)

// BenchmarkQfinite/new-field-12         	 5000000	       237 ns/op	   4.20 MB/s	      80 B/op	       3 allocs/op
// BenchmarkQfinite/field-eq-12          	30000000	        48.0 ns/op	  20.82 MB/s	       0 B/op	       0 allocs/op
// BenchmarkQfinite/field-not-eq-12      	50000000	        27.6 ns/op	  36.18 MB/s	       0 B/op	       0 allocs/op
// BenchmarkQfinite/field-add-12         	 3000000	       488 ns/op	   2.05 MB/s	     128 B/op	       4 allocs/op
// BenchmarkQfinite/field-sub-12         	 5000000	       384 ns/op	   2.60 MB/s	      80 B/op	       3 allocs/op
// BenchmarkQfinite/field-mul-12         	 3000000	       486 ns/op	   2.06 MB/s	     128 B/op	       4 allocs/op
// BenchmarkQfinite/field-pow-12         	 1000000	      1030 ns/op	   0.97 MB/s	     208 B/op	       9 allocs/op
func BenchmarkQfinite(b *testing.B) {
	b.Run("new-field", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = NewQField(2, 31)
		}
	})
	b.Run("field-eq", func(b *testing.B) {
		// initialize
		n1 := QField{*big.NewInt(15), *big.NewInt(31)}
		n2 := QField{*big.NewInt(15), *big.NewInt(31)}

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = n1.Equal(n2)
		}
	})
	b.Run("field-not-eq", func(b *testing.B) {
		// initialize
		n1 := QField{*big.NewInt(15), *big.NewInt(31)}
		n2 := QField{*big.NewInt(7), *big.NewInt(31)}

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = n1.NotEqual(n2)
		}
	})
	b.Run("field-add", func(b *testing.B) {
		// initialize
		n1 := QField{*big.NewInt(17), *big.NewInt(31)}
		n2 := QField{*big.NewInt(21), *big.NewInt(31)}

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = Add(n1, n2)
		}
	})
	b.Run("field-sub", func(b *testing.B) {
		// initialize
		n1 := QField{*big.NewInt(29), *big.NewInt(31)}
		n2 := QField{*big.NewInt(4), *big.NewInt(31)}

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = Sub(n1, n2)
		}
	})
	b.Run("field-mul", func(b *testing.B) {
		// initialize
		n1 := QField{*big.NewInt(29), *big.NewInt(31)}
		n2 := QField{*big.NewInt(4), *big.NewInt(31)}

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = Mul(n1, n2)
		}
	})
	b.Run("field-pow", func(b *testing.B) {
		// initialize
		n1 := QField{*big.NewInt(29), *big.NewInt(31)}

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = Pow(n1, 3)
		}
	})
}
