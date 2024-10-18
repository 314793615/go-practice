package algorithm


import (
	"container/heap"
)

type IntHeap []int

// var _ sort.Interface = (IntHeap)(*nil)


func (h *IntHeap) Less(i, j int) bool{
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Len() int{
	return len(*h)
}

func (h *IntHeap) Swap(i, j int){
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}


func (h *IntHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}


func (h *IntHeap) Pop() interface{}{
	old := *h
	n := len(*h)
	temp := old[n-1]
	*h = old[0:n-1]
	return temp
}


func GetTopN(nums []int, n int) []int {
	if n >= len(nums){
		return nums
	}
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range(nums){
		heap.Push(h, num)
	}
	res := []int{}
	for i:=0;i<n;i++{
		res = append(res, heap.Pop(h).(int))
	}
	return res
}