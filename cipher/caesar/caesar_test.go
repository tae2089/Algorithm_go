package caesar

import "testing"

func TestEncrypt(t *testing.T) {

	tests := []struct {
		name  string
		input string
		key   int
		want  string
	}{
		// TODO: Add test cases.

		{
			"Basic caesar encryption with letter 'a'",
			"a",
			3,
			"d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.input, tt.key); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyEncrypt(t *testing.T) {

	tests := []struct {
		name  string
		input string
		key   int
		want  string
	}{
		// TODO: Add test cases.
		{
			"Basic caesar encryption with letter 'a'",
			"ab",
			3,
			"de",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyEncrypt(tt.input, tt.key); got != tt.want {
				t.Errorf("MyEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDecrypt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   int
		want  string
	}{
		// TODO: Add test cases.
		{
			"Basic caesar decryption with letter 'de'",
			"de",
			3,
			"ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyDecrypt(tt.input, tt.key); got != tt.want {
				t.Errorf("MyDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
