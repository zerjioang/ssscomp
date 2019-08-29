package common

import "testing"

var s IntShare
var sptr *IntShare

// BenchmarkShare/instantiate-12         	2000000000	         1.58 ns/op	 631.73 MB/s	       0 B/op	       0 allocs/op
// BenchmarkShare/instantiate-ptr-12     	50000000	        35.4 ns/op	  28.28 MB/s	       8 B/op	       1 allocs/op
func BenchmarkShare(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			s = NewIntShare(0)
		}
	})
	b.Run("instantiate-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			sptr = NewIntSharePtr(0)
		}
	})
	b.Run("share-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		s = NewIntShare(0)
		for n := 0; n < b.N; n++ {
			_ = s.String()
		}
	})
	b.Run("share-ptr-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		sptr = NewIntSharePtr(0)
		for n := 0; n < b.N; n++ {
			_ = s.String()
		}
	})
}
