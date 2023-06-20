package stringutils

import "testing"

func TestRev(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 1",
			s:    "ABC",
			want: "CBA",
		},
		{
			name: "test 1",
			s:    "АБВ",
			want: "ВБА",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rev(tt.s); got != tt.want {
				t.Errorf("Rev() = %v, want %v", got, tt.want)
			}
		})
	}
}
