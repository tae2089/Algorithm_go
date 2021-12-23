package diffiehellman

import "testing"

func TestGenerateShareKey(t *testing.T) {
	type args struct {
		prvKey int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateShareKey(tt.args.prvKey); got != tt.want {
				t.Errorf("GenerateShareKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateMutualKey(t *testing.T) {
	type args struct {
		prvKey   int64
		shareKey int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMutualKey(tt.args.prvKey, tt.args.shareKey); got != tt.want {
				t.Errorf("GenerateMutualKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
