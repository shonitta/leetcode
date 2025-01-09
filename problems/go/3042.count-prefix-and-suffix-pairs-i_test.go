package leetcode

import (
	"testing"
)

func TestCountPrefixSuffixPairs(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{[]string{"a", "ab", "abc", "abcd"}, 0},
		{[]string{"pa", "papa", "ma", "mama"}, 2},
		{[]string{"abab", "ab"}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countPrefixSuffixPairs(tt.words); got != tt.want {
				t.Errorf("countPrefixSuffixPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
