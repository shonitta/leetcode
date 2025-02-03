package leetcode

import "testing"

func TestLongestMonotonicSubarray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{33}, 1},
		{[]int{1, 4, 3, 3, 2}, 2},
		{[]int{3, 3, 3, 3}, 1},
		{[]int{3, 2, 1}, 3},
	}

	for _, tt := range tests {
		result := longestMonotonicSubarray(tt.input)
		if result != tt.expected {
			t.Errorf("input: %v, got %d, want %d", tt.input, result, tt.expected)
		}
	}
}
