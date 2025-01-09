package leetcode

import (
	"testing"
)

func TestPrefixCount(t *testing.T) {
	tests := []struct {
		words []string
		pref  string
		want  int
	}{
		{[]string{"apple", "app", "apricot", "banana"}, "ap", 3},
		{[]string{"hello", "world", "hell", "heaven"}, "he", 3},
		{[]string{"test", "testing", "tester", "tested"}, "test", 4},
		{[]string{"one", "two", "three"}, "four", 0},
		{[]string{"one", "one", "one"}, "one", 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := prefixCount(tt.words, tt.pref); got != tt.want {
				t.Errorf("prefixCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
