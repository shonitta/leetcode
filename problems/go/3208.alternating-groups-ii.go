/*
* @lc app=leetcode id=3208 lang=golang
*
* [3208] Alternating Groups II
 */

package leetcode

// @lc code=start
// Coded by Cline
// numberOfAlternatingGroups returns the number of alternating groups of length k.
// An alternating group is k contiguous tiles with alternating colors.
func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)

	// Constraints:
	// 3 <= colors.length <= 10^5
	// 0 <= colors[i] <= 1
	// 3 <= k <= colors.length

	// Optimization: Precompute alternating patterns
	// For each position, check if it can start an alternating sequence

	// For a window to be alternating:
	// 1. All adjacent elements must have different colors
	// 2. The window must contain both 0s and 1s

	// We'll use a sliding window approach to efficiently check all windows

	// First, check if the first k-1 positions follow an alternating pattern
	isAlternating := true
	for i := 0; i < k-1; i++ {
		if colors[i] == colors[(i+1)%n] {
			isAlternating = false
			break
		}
	}

	// Count 0s and 1s in the first window
	zeroCount, oneCount := 0, 0
	for i := 0; i < k; i++ {
		if colors[i] == 0 {
			zeroCount++
		} else {
			oneCount++
		}
	}

	// Initialize count
	count := 0
	if isAlternating && zeroCount > 0 && oneCount > 0 {
		count++
	}

	// Sliding window for the rest of the positions
	for start := 1; start < n; start++ {
		// Update the alternating check for the new window
		// Remove the check for the element that's leaving the window
		prevIdx := start - 1
		if colors[prevIdx] != colors[(prevIdx+1)%n] {
			// If the leaving pair was alternating, we need to check the new pair
			newPairStart := (start + k - 2) % n
			newPairEnd := (start + k - 1) % n

			if colors[newPairStart] == colors[newPairEnd] {
				isAlternating = false
			}
		} else {
			// If the leaving pair was not alternating, check if the new window can be alternating
			isAlternating = true
			for i := 0; i < k-1; i++ {
				currIdx := (start + i) % n
				nextIdx := (start + i + 1) % n
				if colors[currIdx] == colors[nextIdx] {
					isAlternating = false
					break
				}
			}
		}

		// Update counts for the sliding window
		if colors[prevIdx] == 0 {
			zeroCount--
		} else {
			oneCount--
		}

		newIdx := (start + k - 1) % n
		if colors[newIdx] == 0 {
			zeroCount++
		} else {
			oneCount++
		}

		// Check if current window is valid
		if isAlternating && zeroCount > 0 && oneCount > 0 {
			count++
		}
	}

	return count
}

// @lc code=end
