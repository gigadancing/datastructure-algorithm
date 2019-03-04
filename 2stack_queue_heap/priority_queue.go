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

func (mp *MinHeap) Top() interface{} {
	return (*mp)[0]
}

// 实现自定义的大顶堆
type MaxHeap []int

func (mp MaxHeap) Len() int {
	return len(mp)
}

func (mp MaxHeap) Less(i, j int) bool {
	return mp[i] > mp[j]
}

func (mp MaxHeap) Swap(i, j int) {
	mp[i], mp[j] = mp[j], mp[i]
}

func (mp *MaxHeap) Pop() interface{} {
	old := *mp
	n := len(old)
	x := old[n-1]
	*mp = old[0 : n-1]
	return x
}

func (mp *MaxHeap) Push(v interface{}) {
	*mp = append(*mp, v.(int))
}

func (mp *MaxHeap) Top() interface{} {
	return (*mp)[0]
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

// 设计一个数据结构，该数据结构动态维护一组数据，且支持如下操作：
// 1. 添加元素：addNum(num int)，将整数num添加至数据结构中。
// 2. 返回数据的中位数: findMedian() float64，返回其维护的数据的中位数。
// 中位数定义：
// 若数据个数为奇数，中位数是该组数据排序后中间的数。
// 若数据个数为偶数，中位数是该组数据排序后中间的两个数字的平均值。
// 思路：
// 用一个大顶堆和一个小顶堆分别存一半数据，两个堆的个数差不超过1；
// 大顶堆堆顶元素小于小顶堆堆顶元素（即大顶堆所有元素小于小顶堆所有元素）。
type MyFind struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func NewMyFind() *MyFind {
	return &MyFind{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (mf *MyFind) addNum(v interface{}) {
	maxHeapLen := mf.maxHeap.Len()
	MinHeapLen := mf.minHeap.Len()

	// 先插入大顶堆
	if maxHeapLen == 0 {
		heap.Push(mf.maxHeap, v)
		return
	}

	if maxHeapLen == MinHeapLen { // 大顶堆和小顶堆的个数相等
		if v.(int) < mf.maxHeap.Top().(int) { // 新元素小于大顶堆堆顶元素，直接插入大顶堆
			heap.Push(mf.maxHeap, v)
		} else { // 插入小顶堆
			heap.Push(mf.minHeap, v)
		}

	} else if maxHeapLen > MinHeapLen { // 大顶堆个数大于小顶堆
		if v.(int) < mf.maxHeap.Top().(int) {
			// 弹出大顶堆堆顶元素插入小顶堆
			val := heap.Pop(mf.maxHeap)
			heap.Push(mf.minHeap, val)
			// 新元素插入大顶堆
			heap.Push(mf.maxHeap, v)
		} else { // 插入小顶堆
			heap.Push(mf.minHeap, v)
		}
	} else { // 大顶堆个数小于小顶堆
		if v.(int) < mf.minHeap.Top().(int) {
			// 插入大顶堆
			heap.Push(mf.maxHeap, v)
		} else {
			// 弹出小顶堆堆顶元素，插入大顶堆
			val := heap.Pop(mf.minHeap)
			heap.Push(mf.maxHeap, val)
			// 新元素插入小顶堆
			heap.Push(mf.minHeap, v)
		}
	}
}

func (mf *MyFind) findMedian() interface{} {
	maxHeapLen := mf.maxHeap.Len()
	MinHeapLen := mf.minHeap.Len()
	if maxHeapLen == MinHeapLen {
		return float64(mf.minHeap.Pop().(int)+mf.maxHeap.Pop().(int)) / 2
	} else if maxHeapLen > MinHeapLen {
		return mf.maxHeap.Pop()
	} else {
		return mf.minHeap.Pop()
	}
	return nil
}
