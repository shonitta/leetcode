package leetcode

import (
	"testing"
)

func TestFindDifferentBinaryString(t *testing.T) {
	tests := []struct {
		name string
		nums []string
		want string
	}{
		{
			name: "Example 1",
			nums: []string{"01", "10"},
			want: "00", // or "11" would also be valid
		},
		{
			name: "Example 2",
			nums: []string{"00", "01"},
			want: "10", // or "11" would also be valid
		},
		{
			name: "Example 3",
			nums: []string{"111", "011", "001"},
			want: "000", // or "010", "100", "101", "110" would also be valid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifferentBinaryString(tt.nums)

			// Check if the result has the same length as input strings
			if len(got) != len(tt.nums[0]) {
				t.Errorf("findDifferentBinaryString() returned string of length %d, want %d",
					len(got), len(tt.nums[0]))
			}

			// Check if the result is a valid binary string
			for _, ch := range got {
				if ch != '0' && ch != '1' {
					t.Errorf("findDifferentBinaryString() returned invalid binary string: %s", got)
				}
			}

			// Check if the result is different from all input strings
			for _, num := range tt.nums {
				if got == num {
					t.Errorf("findDifferentBinaryString() = %v, but result should be different from all input strings", got)
				}
			}
		})
	}
}
