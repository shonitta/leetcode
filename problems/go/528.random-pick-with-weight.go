/*
 * @lc app=leetcode id=528 lang=golang
 *
 * [528] Random Pick with Weight
 */

package leetcode

import (
	"math/rand"
)

// @lc code=start
type Solution struct {
	values   []int
	maxValue int
}

func Constructor(w []int) Solution {
	values := []int{}
	current := 0
	for _, v := range w {
		current += v
		values = append(values, current)
	}
	return Solution{
		values:   values,
		maxValue: current,
	}
}

func (s *Solution) PickIndex() int {
	num := rand.Intn(s.maxValue)
	l, r := 0, len(s.values)-1
	for l <= r {
		m := l + (r-l)/2
		if s.values[m] > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end
