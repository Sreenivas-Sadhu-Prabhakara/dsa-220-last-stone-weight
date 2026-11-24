package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// LastStoneWeight simulates smashing the two heaviest stones together —
// equal stones annihilate, unequal ones leave the difference — and returns
// the final stone's weight (0 when none survives).
func LastStoneWeight(stones []int) int {
	h := maxHeap(append([]int{}, stones...))
	heap.Init(&h)
	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, a-b) // a >= b by pop order
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}

func main() {
	fmt.Println(LastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
