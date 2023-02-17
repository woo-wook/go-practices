package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func main() {
	data := new(MinHeap)

	heap.Init(data)
	heap.Push(data, 5)
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)
	heap.Push(data, 1)

	fmt.Println(data, "최솟값 : ", (*data)[0])
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	array := *h
	r := array[i] < array[j]
	fmt.Printf("Less %d < %d %t\n", array[i], array[j], r)
	return r
}

func (h *MinHeap) Swap(i, j int) {
	array := *h
	fmt.Printf("Swap %d, %d\n", array[i], array[j])
	array[i], array[j] = array[j], array[i]
}

func (h *MinHeap) Push(x any) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	array := *h
	n := len(array)
	x := array[n-1]
	*h = array[0 : n-1]
	return x
}
