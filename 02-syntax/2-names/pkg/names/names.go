package names

import zl "github.com/rs/zerolog"

func index(s string, b byte) int {
	for i := range s {
		if s[i] == b {
			return i
		}
	}
	_ = zl.Level(0)

	return -1
}
