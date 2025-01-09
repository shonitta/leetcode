/*
 * @lc app=leetcode id=2185 lang=golang
 *
 * [2185] Counting Words With a Given Prefix
 */

package leetcode

// @lc code=start
func prefixCount(words []string, pref string) int {
	len_pref := len(pref)
	count := 0
	for _, word := range words {
		if len(word) < len_pref {
			continue
		}

		if word[:len_pref] == pref {
			count++
		}
	}
	return count
}

// @lc code=end
