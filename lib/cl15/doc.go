package main

/*
initial example code performance

BenchmarkCL15/example-12         	       1	3214017795 ns/op	   0.00 MB/s	16631920 B/op	   27663 allocs/op


After removing code std.out logging features:
BenchmarkCL15/example-12         	       1	2248028832 ns/op	   0.00 MB/s	12682072 B/op	   21219 allocs/op

After setting as 'constant' big number e bigE65537
BenchmarkCL15/example-12         	       1	1313508525 ns/op	   0.00 MB/s	 8148656 B/op	   13900 allocs/op

After setting bigOne as constant number
BenchmarkCL15/example-12         	       2	1483222087 ns/op	   0.00 MB/s	 9059852 B/op	   15353 allocs/op

 */