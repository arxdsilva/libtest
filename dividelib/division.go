package dividelib

// Divide returns the result of dividing a by b.
//
// It is the caller's responsibility to ensure that b is not zero.
func Divide(a, b, c int) int {
	return (a + b) / c
}
