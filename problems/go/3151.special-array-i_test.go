package leetcode

import "testing"

func TestIsArraySpecial(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1}, true},
		{[]int{2, 1, 4}, true},
		{[]int{4, 3, 1, 6}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isArraySpecial(tt.nums); got != tt.want {
				t.Errorf("isArraySpecial(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
