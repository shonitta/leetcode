package leetcode

import "testing"

func TestMinimizeXor(t *testing.T) {
	tests := []struct {
		num1 int
		num2 int
		want int
	}{
		{3, 5, 3},
		{1, 12, 3},
		{53, 53, 53},
		{89, 1029, 88},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minimizeXor(tt.num1, tt.num2); got != tt.want {
				t.Errorf("minimizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
