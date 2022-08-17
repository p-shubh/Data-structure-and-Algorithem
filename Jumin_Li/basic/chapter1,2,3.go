package main

import (
	"fmt"
)

// Maxaheap Struct has a slice that holds the array
type Maxaheap struct {
	Array []int
}

// Insert adds an element to the heap
func (h *Maxaheap) Insert(key int) {
	h.Array = append(h.Array, key)
	h.MaxaHeapifyup(len(h.Array) - 1)
}

//Extract return the largest key, and removes it from the heap

func (h *Maxaheap) MaxaHeapifyup(index int) {
	for h.Array[Parent(index)] < h.Array[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
	}
}

//Extract return the largest key, and removes it from the heap

func (h *Maxaheap) Extract() int {
	extracted := h.Array[0]

	l := len(h.Array) - 1

	// when the array is empty
	if len(h.Array) == 0 {
		fmt.Println("cannot extract because the length if the array is 0")
		return -1
	}

	// take out the last index and put it in the root index
	h.Array[0] = h.Array[l]
	h.Array = h.Array[:l]

	h.MaxaHeapifyDown(0)

	return extracted
}

// maxheapify will heapify from bottom top

// maxheapifyDown will be heapify top to bottam
func (h *Maxaheap) MaxaHeapifyDown(index int) {

	lastIndex := len(h.Array) - 1
	l, r := Left(index), Right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.Array[l] > h.Array[r] { // when left child is larger
			childToCompare = l
		} else { //when  right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.Array[index] < h.Array[childToCompare] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = Left(index), Right(index)
		} else {
			return
		}
	}
}

// get the parent index
func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*i + 2
}

// swap key in the array
func (h *Maxaheap) Swap(i1, i2 int) {
	h.Array[i1], h.Array[i2] = h.Array[i2], h.Array[i1]
}

func main() {
	m := &Maxaheap{}
	fmt.Println(m)
	buidHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range buidHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	// for i := 0; i < 4; i++ {
	// 	m.Extract()
	// 	fmt.Println(m)
	// }
}
