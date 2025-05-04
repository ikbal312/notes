// source: https://github.com/golang/go/blob/master/src/container/heap/heap.go
// heap or priority queue implementation in Go
// mean-heap or minimun element at the root
package main

import "sort"

/*
Heap is a complete binary tree that satisfies the heap property.

Heap Property:
  - Min Heap: The value of each node is greater than or equal to the value of its parent, with the minimum-value element at the root.
  - Max Heap: The value of each node is less than or equal to the value of its parent, with the maximum-value element at the root.

used In:
  - Priority Queue
*/
type Interface interface {
	sort.Interface
	Push(x any)
	Pop() any
}

func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- { // start from the last parent node
		down(h, i, n) // heapify
	}
}

// add x to the heap
func Push(h Interface, x any) {
	h.Push(x)        // append x to the end of the heap
	up(h, h.Len()-1) // move x up to the correct position

}

func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)   // swap the root with the last element
	down(h, 0, n)  // move the new root down to the correct position
	return h.Pop() // return the last element
}

// Remove removes the element at index i from the heap.
func Remove(h Interface, i int) any {
	n := h.Len() - 1 // last index
	if n != i {
		h.Swap(i, n)        // swap the element to remove with the last element
		if !down(h, i, n) { // if the element is not moved down
			up(h, i) // move the element up
		}
	}
	return h.Pop() // return the last element
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(h Interface, j int) { // j is the index of the element to move up
	for {
		i := (j - 1) / 2             // parent
		if i == j || !h.Less(j, i) { // if j is the root or parent >= child
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h Interface, i0, n int) bool {
	i := i0 // i is the parent
	for {
		j1 := 2*i + 1          // left child
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}

		j := j1                                     // j = left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) { // j2 = right child
			j = j2 // j = right child
		}

		if !h.Less(j, i) { // if parent >= child
			break
		}

		h.Swap(i, j) // swap parent and child
		i = j        // move down
	}
	return i > i0 // return true if moved down
}
