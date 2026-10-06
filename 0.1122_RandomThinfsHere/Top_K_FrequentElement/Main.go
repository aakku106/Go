package topkfrequentelement

import "fmt"

func topKFrequent(nums []int, k int) []int {
	g := make(map[int]int)

	for _, v := range nums {
		g[v]++
	}

	for i, v := range g {
		fmt.Println(i, v)
	}

	return nil
}
