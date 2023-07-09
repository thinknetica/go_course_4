package table

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Тест №1",
			args: args{
				s: "string",
			},
			want: "gnirts",
		},
		{
			name: "Тест №2",
			args: args{
				s: "ABCdef",
			},
			want: "fedCBA",
		},
		{
			name: "Тест №3",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
