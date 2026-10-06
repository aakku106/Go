package topkfrequentelement

import (
	"cmp"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	group := make(map[int]int)

	for _, v := range nums {
		group[v]++
	}

	var bucket [][]int
	for key, value := range group {
		bucket = append(bucket, []int{key, value})
	}
	slices.SortFunc(bucket, func(first, second []int) int {
		return cmp.Compare(second[1], first[1])
	})
	var final []int
	for i := range k {
		final = append(final, bucket[i][0])
	}
	return final
}
