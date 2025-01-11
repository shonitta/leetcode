package leetcode

import (
	"testing"
)

func TestWordSubsets(t *testing.T) {
	tests := []struct {
		words1 []string
		words2 []string
		want   []string
	}{
		// Return some words
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}, []string{"facebook", "google", "leetcode"}},
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}, []string{"apple", "google", "leetcode"}},
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}, []string{"facebook", "google"}},
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"pep"}, []string{"apple"}},
		// Return no words
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "zzzzz"}, []string{}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := wordSubsets(tt.words1, tt.words2); !equal(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
