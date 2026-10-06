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

	bucket := make([][2]int, 0, len(group))
	for key, value := range group {
		bucket = append(bucket, [2]int{key, value})
	}
	slices.SortFunc(bucket, func(first, second [2]int) int {
		return cmp.Compare(second[1], first[1])
	})
	final := make([]int, 0, k)
	for i := range k {
		final = append(final, bucket[i][0])
	}
	return final
}
