package main

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	listItems := []*Item{
		{Name: "Carrot", Expiry: 30},
		{Name: "Potato", Expiry: 45},
		{Name: "Rice", Expiry: 100},
		{Name: "Spinach", Expiry: 5},
	}

	priorityQueue := make(PriorityQueue, len(listItems))

	for i, item := range listItems {
		priorityQueue[i] = item
		priorityQueue[i].Index = i
	}

	heap.Init(&priorityQueue)

	// Print the order by Priority of expiry
	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		fmt.Printf("Name: %s Expiry: %d\n", item.Name, item.Expiry)
	}
}

func TestMinHeap_FindKthLargest(t *testing.T) {
	array := []int{3, 2, 1, 5, 6, 4}
	kthMax := FindKthLargest(array, 2)
	fmt.Println(kthMax)
}

func TestMinHeap(t *testing.T) {
	mp := &MinHeap{}
	heap.Init(mp)

	heap.Push(mp, 3)
	fmt.Println(mp.Top()) // 3

	heap.Push(mp, 1)
	fmt.Println(mp.Top()) // 1

	heap.Push(mp, 2)
	fmt.Println(mp.Top()) // 1

	heap.Push(mp, 5)
	fmt.Println(mp.Top()) // 1

	heap.Push(mp, 6)
	fmt.Println(mp.Top()) // 1

	heap.Push(mp, 4)
	fmt.Println(mp.Top()) // 1

}

func TestMyFind(t *testing.T) {
	mf := NewMyFind()
	mf.addNum(3)
	mf.addNum(1)
	mf.addNum(2)
	mf.addNum(5)
	mf.addNum(6)
	mf.addNum(4)
	fmt.Println(mf.findMedian())
}
