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
			hasPrefix := strings.HasPrefix(words[j], words[i])
			hasSuffix := strings.HasSuffix(words[j], words[i])
			if hasPrefix && hasSuffix {
				count++
			}
		}
	}
	return count
}

// @lc code=end
