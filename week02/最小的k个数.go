package week02

import "container/heap"

//minHeap
func getLeastNumbers(arr []int, k int) []int {
	res := make([]int, 0, k)
	h := &minIntQueue{}
	heap.Init(h)
	for _, v := range arr {
		heap.Push(h, v)
	}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

type minIntQueue []int

func (h minIntQueue) Len() int           { return len(h) }
func (h minIntQueue) Less(i, j int) bool { return h[i] < h[j] }
func (h minIntQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minIntQueue) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minIntQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

//maxHeap
func getLeastNumbers(arr []int, k int) []int {
	res := make([]int, 0, k)
	h := &minIntQueue{}
	heap.Init(h)
	for _, v := range arr {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

type minIntQueue []int

func (h minIntQueue) Len() int           { return len(h) }
func (h minIntQueue) Less(i, j int) bool { return h[i] > h[j] }
func (h minIntQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minIntQueue) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minIntQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
