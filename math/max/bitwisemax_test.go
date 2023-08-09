package max

import "testing"

func TestBitwise(t *testing.T) {
	t.Run("Testing(64bit) a = 32 and m = 9223372036854775807: ", func(t *testing.T) {
		max := Bitwise(32, 9223372036854775807)
		if max != 9223372036854775807 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(64bit) a = 1024 and m = -9223372036854770001: ", func(t *testing.T) {
		max := Bitwise(1024, -9223372036854770001)
		if max != 1024 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(64bit) a = -6 and m = -6: ", func(t *testing.T) {
		max := Bitwise(-6, -6)
		if max != -6 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})

	t.Run("Testing(64bit) a = -4223372036854775809 and m = -4223372036854775808: ", func(t *testing.T) {
		max := Bitwise(-4223372036854775809, -4223372036854775808)
		if max != -4223372036854775808 {
			t.Fatalf("Error: Bitwise returned bad value")
		}
	})
}
