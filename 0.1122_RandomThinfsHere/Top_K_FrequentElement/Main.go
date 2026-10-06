package topkfrequentelement

func topKFrequent(nums []int, k int) []int {
	g := make(map[int][][2]int)

	for _, v := range nums {
		if _, ok := g[v]; ok {
			g[v][0][1] += 1
		}
	}

	return nil
}
