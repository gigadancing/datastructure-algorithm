package main

import (
	"container/heap"
)

type Item struct {
	Name   string
	Expiry int
	Index  int
}

type PriorityQueue []*Item

// sort

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Expiry < pq[j].Expiry
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// heap

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	item.Index = -1

	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// 实现自定义的小顶堆
type MinHeap []int

func (mp MinHeap) Len() int {
	return len(mp)
}

func (mp MinHeap) Less(i, j int) bool {
	return mp[i] < mp[j]
}

func (mp MinHeap) Swap(i, j int) {
	mp[i], mp[j] = mp[j], mp[i]
}

func (mp *MinHeap) Pop() interface{} {
	old := *mp
	n := len(old)
	x := old[n-1]
	*mp = old[0 : n-1]
	return x
}

func (mp *MinHeap) Push(v interface{}) {
	*mp = append(*mp, v.(int))
}

// 一个未排序的数组，求这个数组中第K大的数
// 如 array = [3, 2, 1, 5, 6, 4], k = 2, return 5
// 思路：
// 维护一个K大小的小顶堆，堆中元素个数小于K时，直接加入新元素；否则，和栈顶元素比较，如果大于栈顶元素，则弹出栈顶元素，加入新元素
func FindKthLargest(array []int, k int) int {
	if len(array) == 0 || k >= len(array) || k <= 0 {
		return -1
	}
	mp := &MinHeap{}
	heap.Init(mp)

	for i := 0; i < len(array); i++ {
		if i+1 <= k {
			heap.Push(mp, array[i])
		} else {
			v := heap.Pop(mp).(int)
			if v < array[i] {
				heap.Push(mp, array[i])

			} else {
				heap.Push(mp, v)
			}

		}
	}

	return heap.Pop(mp).(int)
}
