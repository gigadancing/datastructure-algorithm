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
