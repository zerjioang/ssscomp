package ed25519

import "testing"

func BenchmarkInternals(b *testing.B) {
	b.Run("zero-slice", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		elem := fieldElement{}
		for n := 0; n < b.N; n++ {
			FeZero(&elem)
		}
	})
	b.Run("one-slice", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		elem := fieldElement{}
		for n := 0; n < b.N; n++ {
			FeOne(&elem)
		}
	})
	b.Run("add-fe", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		a := fieldElement{2,2,2,2,2,2,2,2,2,2}
		c := fieldElement{3,3,3,3,3,3,3,3,3,3}
		dst := fieldElement{}
		for n := 0; n < b.N; n++ {
			FeAdd(&dst, &a, &c)
		}
	})
}
