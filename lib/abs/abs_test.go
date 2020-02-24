/*
Package abs provides multiple implementations of the abs function, to compute
the absolute value of a signed, 64-bit integer.

This package complements the following article, which compares each
implementation:

http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
*/

package abs_test

import (
	"math"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/zerjioang/ssscomp/lib/abs"
)

// WithBranch uses control structures to return the absolute value of an
// integer.
func WithBranch(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// WithStdLib uses the standard library's math package to compute the
// absolute value on an integer.
//
// We expect test for correctness to fail on large numbers that overflow
// float64.
func WithStdLib(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

// WithTwosComplement uses a trick from Henry S. Warren's incredible book,
// Hacker's Delight. It utilizes Two's Complement arithmetic to compute the
// absolute value of an integer.
func WithTwosComplement(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

const (
	MaxInt int64 = 1<<63 - 1
	MinInt int64 = -1 << 63
)

// An absFunc is a function that returns the absolute value of an integer.
type absFunc func(int64) int64

func funcName(v interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	return s[strings.LastIndex(s, ".")+1:]
}

func TestAbs(t *testing.T) {
	inputs := []int64{MinInt + 1, MinInt + 2, -1, -0, 1, 2, MaxInt - 1, MaxInt}
	outputs := []int64{MaxInt, MaxInt - 1, 1, 0, 1, 2, MaxInt - 1, MaxInt}
	testFuncs := []absFunc{
		WithBranch,
		// WithStdLib, // test failure expected on large numbers
		WithTwosComplement,
		abs.WithASM,
	}
	for _, f := range testFuncs {
		testName := funcName(f)
		t.Run(testName, func(t *testing.T) {
			for i := 0; i < len(inputs); i++ {
				actual := f(inputs[i])
				if actual != outputs[i] {
					t.Errorf("%s(%d)", testName, inputs[i])
					t.Errorf("	input:		%064b (%d)", uint64(inputs[i]), inputs[i])
					t.Errorf("	expected:	%064b (%d)", uint64(outputs[i]), outputs[i])
					t.Errorf("	actual:		%064b (%d)", uint64(actual), actual)
				}
			}
		})
	}
}

var r uint64 = 0xdeadbeef

// Pseudo-random number generator adapted from
// https://github.com/dgryski/trifles/blob/master/fastrand/fastrand.go
func Rand() int64 {
	r ^= r >> 12 // a
	r ^= r << 25 // b
	r ^= r >> 27 // c
	return int64(r * 2685821657736338717)
}
