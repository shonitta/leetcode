package leetcode

import (
	"testing"
)

func TestRepairCars(t *testing.T) {
	testCases := []struct {
		name     string
		ranks    []int
		cars     int
		expected int64
	}{
		{
			name:     "Example 1",
			ranks:    []int{4, 2, 3, 1},
			cars:     10,
			expected: 16,
		},
		{
			name:     "Example 2",
			ranks:    []int{5, 1, 8},
			cars:     6,
			expected: 16,
		},
		{
			name:     "Single mechanic",
			ranks:    []int{3},
			cars:     10,
			expected: 300, // 3 * 10^2
		},
		{
			name:     "Same rank mechanics",
			ranks:    []int{2, 2, 2, 2},
			cars:     16,
			expected: 32, // Each mechanic can repair 4 cars in 32 minutes (2*4^2=32), so 4 mechanics can repair 16 cars
		},
		{
			name:     "Large number of cars",
			ranks:    []int{1, 2, 3, 4, 5},
			cars:     100,
			expected: 980, // Optimal distribution among mechanics leads to this time
		},
		{
			name:     "Edge case - one car",
			ranks:    []int{10, 5, 3},
			cars:     1,
			expected: 3, // The mechanic with rank 3 can repair 1 car in 3 minutes
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := repairCars(tc.ranks, tc.cars)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
