package paillier

/*
Paillier 2 implementation benchmarking

BenchmarkKey1024-12                	      20	  58871168 ns/op
BenchmarkKey2048-12                	       2	 734826996 ns/op
BenchmarkKey3072-12                	       1	3021790128 ns/op
BenchmarkKey4096-12                	       1	3704751256 ns/op
BenchmarkEncryptionSmall1024-12    	     300	   5611977 ns/op
BenchmarkEncryptionSmall2048-12    	      50	  38566454 ns/op
BenchmarkEncryptionSmall3072-12    	      10	 124464305 ns/op
BenchmarkEncryptionSmall4096-12    	       5	 282421117 ns/op
BenchmarkEncryptionLarge1024-12    	     300	   5814865 ns/op
BenchmarkEncryptionLarge2048-12    	      50	  38576939 ns/op
BenchmarkEncryptionLarge3072-12    	      10	 123609317 ns/op
BenchmarkEncryptionLarge4096-12    	       5	 272374324 ns/op
BenchmarkDecryptionSmall1024-12    	    1000	   1979709 ns/op
BenchmarkDecryptionSmall2048-12    	     100	  11946544 ns/op
BenchmarkDecryptionSmall3072-12    	      30	  40615229 ns/op
BenchmarkDecryptionSmall4096-12    	      10	 104521368 ns/op
BenchmarkDecryptionLarge1024-12    	    1000	   2002632 ns/op
BenchmarkDecryptionLarge2048-12    	     100	  12172811 ns/op
BenchmarkDecryptionLarge3072-12    	      30	  38650074 ns/op
BenchmarkDecryptionLarge4096-12    	      10	 103891719 ns/op
BenchmarkAdditionLarge1024-12      	  100000	     12870 ns/op
BenchmarkAdditionLarge2048-12      	   30000	     37761 ns/op
BenchmarkAdditionLarge3072-12      	   20000	     78113 ns/op
BenchmarkAdditionLarge4096-12      	    5000	    212566 ns/op

*/
