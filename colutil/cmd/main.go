package main

import (
	"container/heap"
	"fmt"

	cu "github.com/dbraley/advent-of-code/colutil"
)

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(cu.PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &cu.Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &cu.Item{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*cu.Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
}
