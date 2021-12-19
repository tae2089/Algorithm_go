package checksum

import "testing"

func TestLuhnAlgorithm(t *testing.T) {

	tests := []struct {
		name string
		s    []rune
		want bool
	}{
		// TODO: Add test cases.
		{"check 4242424242424242", []rune("4242424242424242"), true},
		{"check 6200000000000005", []rune("6200000000000005"), true},
		{"check 5534200028533164", []rune("5534200028533164"), true},
		{"check 36227206271667", []rune("36227206271667"), true},
		{"check 471629309440", []rune("471629309440"), false},
		{"check 1111", []rune("1111"), false},
		{"check 12345674", []rune("12345674"), true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuhnAlgorithm(tt.s); got != tt.want {
				t.Errorf("LuhnAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
