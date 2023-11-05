package taskexecutor

type heapData []processData

func (h heapData) Len() int {
	return len(h)
}

func (h heapData) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h heapData) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapData) Push(x interface{}) {
	*h = append(*h, x.(processData))
}

func (h *heapData) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
