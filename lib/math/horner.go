package math

// Horner's rule for polynomial evaluation
// compute the result from the innermost brackets outwards
// https://en.wikipedia.org/wiki/Horner%27s_method
func Horner(x int64, c []int64) (acc int64) {
	for i := len(c) - 1; i >= 0; i-- {
		acc = acc*x + c[i]
	}
	return
}
