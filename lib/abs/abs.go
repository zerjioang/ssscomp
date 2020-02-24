package abs

// optimized abs function for int64 data type
func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
