package twosum

import ()

func twoSum(nums []int, target int) []int {

	HM := make(map[int]int, len(nums))
	for i := range nums {
		if j, ok := HM[target-nums[i]]; ok {
			return []int{j, i}
		}
		HM[nums[i]] = i
	}

	return nil
}

func twosumBrutal(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[j] == nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}
