package unpaddedrsa

import (
	"math/big"
	"testing"
)

// BenchmarkUnpaddedRSA/generate-key-1024-12         	      50	  60468243 ns/op	   0.00 MB/s	 1135332 B/op	    4292 allocs/op
// BenchmarkUnpaddedRSA/generate-key-2048-12         	       2	 670219801 ns/op	   0.00 MB/s	 3664436 B/op	    9484 allocs/op
// BenchmarkUnpaddedRSA/generate-key-4096-12         	       1	7305552893 ns/op	   0.00 MB/s	13562288 B/op	   22802 allocs/op
// BenchmarkUnpaddedRSA/1024k-encrypt-12             	   30000	     38173 ns/op	   0.03 MB/s	    1865 B/op	      15 allocs/op
// BenchmarkUnpaddedRSA/homomorphic-multiplication-12    2000000 	       900 ns/op	   1.11 MB/s	     288 B/op	       1 allocs/op
func BenchmarkUnpaddedRSA(b *testing.B) {
	b.Run("generate-key-1024", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = GenerateKeyPair(1024)
		}
	})
	b.Run("generate-key-2048", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = GenerateKeyPair(2048)
		}
	})
	b.Run("generate-key-4096", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = GenerateKeyPair(4096)
		}
	})
	b.Run("1024k-encrypt", func(b *testing.B) {
		// key generation
		_, pub := GenerateKeyPair(1024)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			x1 := 200
			b1 := big.NewInt(int64(x1))
			cipher1, err := EncryptUnpaddedRSA(pub, b1)
			if err != nil || cipher1 == nil {
				b.FailNow()
			}
		}
	})
	b.Run("homomorphic-multiplication", func(b *testing.B) {
		// generate key
		_, pub := GenerateKeyPair(1024)
		x1 := 200
		x2 := 3
		// we can compute
		// enc(x1) * enc(x2) == ( x1^e * x2^e ) mod n
		b1 := big.NewInt(int64(x1))
		b2 := big.NewInt(int64(x2))

		// unpadded RSA encrypted version of x1
		cipher1, _ := EncryptUnpaddedRSA(pub, b1)
		// unpadded RSA encrypted version of x2
		cipher2, _ := EncryptUnpaddedRSA(pub, b2)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// compute homomorphic multiplication over encrypted values cipher1 and cipher2
			c12 := big.NewInt(int64(0))
			c12.Mul(cipher1, cipher2)
		}
	})
}
