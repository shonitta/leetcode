/*
 * @lc app=leetcode id=3042 lang=golang
 *
 * [3042] Count Prefix and Suffix Pairs I
 */

// @lc code=start
package leetcode

import "strings"

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			has_prefix := strings.HasPrefix(words[j], words[i])
			has_suffix := strings.HasSuffix(words[j], words[i])
			if has_prefix && has_suffix {
				count++
			}
		}
	}
	return count
}

// @lc code=end
