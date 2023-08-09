package max

const Base64Bit = 63
const Base32Bit = 31

// Bitwise computes using bitwise operator the maximum of all the integer input and returns it
func Bitwise(a int, b int) int {
	z := a - b
	i := (z >> Base64Bit) & 1
	return a - (i * z)
}
