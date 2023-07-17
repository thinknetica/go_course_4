package tdd

import "testing"

func TestFact(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Test #2",
			args: args{n: 3},
			want: 6,
		},
		{
			name: "Test #2",
			args: args{n: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fact(tt.args.n); got != tt.want {
				t.Errorf("Fact() = %v, want %v", got, tt.want)
			}
		})
	}
}
