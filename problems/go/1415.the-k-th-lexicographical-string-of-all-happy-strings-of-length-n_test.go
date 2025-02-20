package leetcode

import "testing"

func Test_getHappyString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{
			name: "n=1, k=3",
			n:    1,
			k:    3,
			want: "c",
		},
		{
			name: "n=1, k=4",
			n:    1,
			k:    4,
			want: "",
		},
		{
			name: "n=3, k=9",
			n:    3,
			k:    9,
			want: "cab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHappyString(tt.n, tt.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
