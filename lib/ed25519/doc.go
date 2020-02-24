package ed25519

/*
This is a fork of Go ed25519 curve implementation and this are initial performance values

BenchmarkScMul-12                  	10368472	       117 ns/op
BenchmarkScMulAdd-12               	 9671659	       120 ns/op
BenchmarkPointMult-12              	   12844	     94594 ns/op
BenchmarkDoublePointMult-12        	   12985	     92362 ns/op
BenchmarkProj2Ext-12               	  154582	      7799 ns/op
BenchmarkProjBytesExt-12           	   70645	     16309 ns/op
BenchmarkInvertModL-12             	   35466	     33748 ns/op
BenchmarkKeyGeneration-12          	   26034	     45948 ns/op	   0.02 MB/s	     128 B/op	       3 allocs/op
BenchmarkSigning-12                	   25561	     48417 ns/op	   0.02 MB/s	     512 B/op	       6 allocs/op
BenchmarkVerification-12           	    8808	    127150 ns/op	   0.01 MB/s	     288 B/op	       2 allocs/op
BenchmarkPublicKeyExtraction-12    	    5078	    237672 ns/op
BenchmarkSigningExt-12             	   23332	     47784 ns/op
BenchmarkVerificationExt-12        	    9492	    125750 ns/op

*/
