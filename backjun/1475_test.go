package backjun

import "testing"

func Test_func1475(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "base1",
			args: args{
				data: "9999",
			},
			want: 2,
		},
		{
			name: "base2",
			args: args{
				data: "122",
			},
			want: 2,
		},
		{
			name: "base3",
			args: args{
				data: "12635",
			},
			want: 1,
		},
		{
			name: "base4",
			args: args{
				data: "888888",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := func1475(tt.args.data); got != tt.want {
				t.Errorf("func1475() = %v, want %v", got, tt.want)
			}
		})
	}
}
