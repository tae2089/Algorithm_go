package diffiehellman

const (
	generator         = 3
	primeNumber int64 = 6700417
)

func GenerateShareKey(prvKey int64) int64 {
	return modularExponentiation(generator, prvKey, primeNumber)
}

func GenerateMutualKey(prvKey, shareKey int64) int64 {
	return modularExponentiation(shareKey, prvKey, primeNumber)
}

func modularExponentiation(b, e, mod int64) int64 {
	if mod == 1 {
		return 0
	}
	var r int64 = 1
	b = b % mod
	for e > 0 {
		if e&1 == 1 {
			r = (r * b) % mod
		}
		e = e >> 1
		b = (b * b) % mod
	}
	return r
}
