package leetcode

import (
	"testing"
)

func TestDivideArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1 - Valid array with pairs",
			nums:     []int{3, 2, 3, 2, 2, 2},
			expected: true,
		},
		{
			name:     "Example 2 - Invalid array without pairs",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Example 3 - Unsorted array with valid pairs",
			nums:     []int{1, 2, 1, 2},
			expected: true,
		},
		{
			name:     "Sorted array with valid pairs",
			nums:     []int{2, 2, 4, 4, 6, 6},
			expected: true,
		},
		{
			name:     "Unsorted array with invalid pairs",
			nums:     []int{5, 7, 3, 2, 1, 8},
			expected: false,
		},
		{
			name:     "Array with same elements",
			nums:     []int{5, 5, 5, 5, 5, 5},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := divideArray(tc.nums)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
