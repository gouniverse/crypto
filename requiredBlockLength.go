package crypto

// RequiredBlockLength calculates the block length
// required (128) to contain the provided length
func RequiredBlockLength(v int) int {
	a := 128
	for v > a {
		a = a * 2
	}
	return a
}
