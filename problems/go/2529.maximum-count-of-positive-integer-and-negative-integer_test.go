package leetcode

import "testing"

func Test_maximumCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1: Equal number of positive and negative integers",
			nums: []int{-2, -1, -1, 1, 2, 3},
			want: 3,
		},
		{
			name: "Example 2: More negative than positive integers with zeros",
			nums: []int{-3, -2, -1, 0, 0, 1, 2},
			want: 3,
		},
		{
			name: "Example 3: Only positive integers",
			nums: []int{5, 20, 66, 1314},
			want: 4,
		},
		{
			name: "Additional Test 1: Only negative integers",
			nums: []int{-10, -9, -8, -7, -6},
			want: 5,
		},
		{
			name: "Additional Test 2: Only zeros",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "Additional Test 3: Single element - positive",
			nums: []int{42},
			want: 1,
		},
		{
			name: "Additional Test 4: Single element - negative",
			nums: []int{-42},
			want: 1,
		},
		{
			name: "Additional Test 5: Single element - zero",
			nums: []int{0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
