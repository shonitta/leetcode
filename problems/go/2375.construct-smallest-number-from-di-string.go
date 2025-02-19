/*
 * @lc app=leetcode id=2375 lang=golang
 *
 * [2375] Construct Smallest Number From DI String
 */

package leetcode

import "strconv"

// @lc code=start
func smallestNumber(pattern string) string {
	n := len(pattern)
	nums := make([]int, n+1)
	for i := range nums {
		nums[i] = i + 1
	}

	for i := 0; i < n; i++ {
		if pattern[i] == 'D' {
			start := i
			for i < n && pattern[i] == 'D' {
				i++
			}
			for left, right := start, i; left < right; left, right = left+1, right-1 {
				nums[left], nums[right] = nums[right], nums[left]
			}
			i--
		}
	}

	result := ""
	for _, num := range nums {
		result += strconv.Itoa(num)
	}
	return result
}

// @lc code=end
