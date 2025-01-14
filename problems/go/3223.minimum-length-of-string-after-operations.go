/*
 * @lc app=leetcode id=3223 lang=golang
 *
 * [3223] Minimum Length of String After Operations
 */

package leetcode

// @lc code=start
func minimumLength(s string) int {
	// TODO: byte map is faster than rune map
	cMap := make(map[rune]int)
	for _, c := range s {
		cMap[c]++
	}
	res := 0
	for _, count := range cMap {
		if count > 2 {
			if count%2 == 0 {
				res += 2
			} else {
				res += 1
			}
		} else {
			res += count
		}
	}
	return res
}

// @lc code=end
