package leetcode

import "testing"

func TestMinimumLength(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abaacbcbb", 5},
		{"aa", 2},
		{"a", 1},
		{"aaaaaaaaaaaaaaaaa", 1},
		{"ababababababababababababababababa", 3},
		{"abcdefghijklmn", 14},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minimumLength(tt.s); got != tt.want {
				t.Errorf("minimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
