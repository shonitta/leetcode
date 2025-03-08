/*
 * @lc app=leetcode id=2379 lang=golang
 *
 * [2379] Minimum Recolors to Get K Consecutive Black Blocks
 */

package leetcode

// @lc code=start
func minimumRecolors(blocks string, k int) int {
	whiteCount := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			whiteCount++
		}
	}

	res := whiteCount
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			whiteCount++
		}
		if blocks[i-k] == 'W' {
			whiteCount--
		}
		res = min(res, whiteCount)
	}
	return res
}

// @lc code=end
