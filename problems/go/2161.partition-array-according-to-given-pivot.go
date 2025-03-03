/*
 * @lc app=leetcode id=2161 lang=golang
 *
 * [2161] Partition Array According to Given Pivot
 */

package leetcode

// @lc code=start
func pivotArray(nums []int, pivot int) []int {
	less := []int{}
	equal := []int{}
	greater := []int{}

	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greater = append(greater, num)
		}
	}

	return append(less, append(equal, greater...)...)
}

// @lc code=end
