package leetcode

import (
	"strings"
	"testing"
)

func TestCanBeValid(t *testing.T) {
	tests := []struct {
		s      string
		locked string
		want   bool
	}{
		// Even Length Cases
		{"()", "11", true},
		{")(", "11", false},
		{"((()))", "111111", true},
		{"((()()", "111101", true},
		{"((()))", "000000", true},                               // No locked parentheses always valid
		{strings.Repeat("(", 10), strings.Repeat("0", 10), true}, // No locked parentheses always valid
		// Odd Length Cases
		{"(", "0", false},
		{"((())", "11110", false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := canBeValid(tt.s, tt.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
