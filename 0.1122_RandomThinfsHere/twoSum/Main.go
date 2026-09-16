package twosum

import ()

func twoSum(nums []int, target int) []int {

	// HM := make(map[int]int, len(nums))
	//
	return nil
}

func twosumBrutal(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if target-nums[j] == nums[i] {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return nil
}
