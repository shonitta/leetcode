/*
 * @lc app=leetcode id=1726 lang=golang
 *
 * [1726] Tuple with Same Product
 */

package leetcode

// @lc code=start
func tupleSameProduct(nums []int) int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			hashMap[nums[i]*nums[j]]++
		}
	}

	res := 0
	for _, v := range hashMap {
		res += v * (v - 1) * 4
	}
	return res
}

// @lc code=end
