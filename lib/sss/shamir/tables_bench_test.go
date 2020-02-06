package shamir

import "testing"

// BenchmarkTableLookup/log-table-first-12         	2000000000	         0.80 ns/op	1255.99 MB/s	       0 B/op	       0 allocs/op
// BenchmarkTableLookup/log-table-last-12          	2000000000	         0.79 ns/op	1264.31 MB/s	       0 B/op	       0 allocs/op
// BenchmarkTableLookup/exp-table-first-12         	2000000000	         1.59 ns/op	 630.05 MB/s	       0 B/op	       0 allocs/op
// BenchmarkTableLookup/exp-table-last-12          	2000000000	         0.80 ns/op	1253.63 MB/s	       0 B/op	       0 allocs/op
func BenchmarkTableLookup(b *testing.B) {
	b.Run("log-table-first", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = logTable[0]
		}
	})
	b.Run("log-table-last", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = logTable[255]
		}
	})
	b.Run("exp-table-first", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = expTable[0]
		}
	})
	b.Run("exp-table-last", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = expTable[255]
		}
	})
}
