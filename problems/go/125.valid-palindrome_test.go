package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"a b; cd: b, a", false},
		{"100aba001", true},
		{"000aba100", false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
