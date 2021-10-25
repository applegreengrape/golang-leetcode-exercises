package heap

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h [j]
}

func (h IntHeap) Swap (i, j int ) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push (x interface{}){
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minHeap(h IntHeap, num int) IntHeap {
	heap.Init(&h)
	heap.Push(&h, num)

	return h
}

type MaxHeap struct {
	IntHeap
}

func (h MaxHeap) Less (i, j int) bool {
	return h.IntHeap[i] > h.IntHeap[j]
}

func maxHeap(h MaxHeap, num int) MaxHeap {
	heap.Init(&h)
	heap.Push(&h, num)

	return h
}