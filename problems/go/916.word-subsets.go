/*
 * @lc app=leetcode id=916 lang=golang
 *
 * [916] Word Subsets
 */

package leetcode

// @lc code=start
func wordSubsets(words1 []string, words2 []string) []string {
	charCount := make([]int, 26)
	for _, word := range words2 {
		count := make([]int, 26)
		for _, c := range word {
			count[c-'a']++
			if count[c-'a'] > charCount[c-'a'] {
				charCount[c-'a'] = count[c-'a']
			}
		}
	}

	res := []string{}
	for _, word := range words1 {
		count := make([]int, 26)
		for _, c := range word {
			count[c-'a']++
		}

		valid := true
		for i := 0; i < 26; i++ {
			if count[i] < charCount[i] {
				valid = false
				break
			}
		}

		if valid {
			res = append(res, word)
		}
	}
	return res
}

// @lc code=end
