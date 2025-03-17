/*
 * @lc app=leetcode id=2594 lang=golang
 *
 * [2594] Minimum Time to Repair Cars
 */

package leetcode

import (
	"math"
)

// @lc code=start
func repairCars(ranks []int, cars int) int64 {
	left, right := int64(1), int64(1e14)

	for left < right {
		mid := left + (right-left)/2

		if canRepairAllCars(ranks, cars, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canRepairAllCars(ranks []int, cars int, time int64) bool {
	totalCarsRepaired := 0

	for _, rank := range ranks {
		carsRepairedByMechanic := int(math.Sqrt(float64(time) / float64(rank)))
		totalCarsRepaired += carsRepairedByMechanic
	}

	return totalCarsRepaired >= cars
}

// @lc code=end
