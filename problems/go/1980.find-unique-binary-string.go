package leetcode

/*
 * @lc app=leetcode id=1980 lang=golang
 *
 * [1980] Find Unique Binary String
 */

// @lc code=start
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	res := ""
	for i := 0; i < n; i++ {
		if nums[i][i] == '0' {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

// @lc code=end
