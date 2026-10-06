package topkfrequentelement

type countTable struct {
	value     int
	frequency int
}

func topKFrequent(nums []int, k int) []int {
	g := make(map[int]countTable)

	for _, v := range nums {
		g[v]={v,}
	}

	return nil
}
