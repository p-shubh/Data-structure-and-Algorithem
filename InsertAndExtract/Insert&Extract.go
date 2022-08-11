package main

import "fmt"

// Maxheap struct has a slice that holds the array
type Maxheap struct {
	Array []int
}

// Insert adds an element to the heap
func (h Maxheap) Insert(key int) {

	h.Array = append(h.Array, key)
	h.MaxHeapifyUp(len(h.Array) - 1)
}

// Extract return the largest key, removes it from the heap

func (h Maxheap) MaxHeapifyUp(index int) {

}

func (h Maxheap) Parent(left_child_index int) int {
	

	Parent := (left_child_index - 1) / 2

	return Parent

}

func Left(h Maxheap) {
	h.MaxHeapifyUp(7)
	h.Parent(7)

}

func main() {
	reqBody := Maxheap{}

	c := reqBody.Parent(8)

	fmt.Println("c", c)

}

// func firstBadVersion(n int) int {

//     i, j := 1,n

// 	for i + 1 < j {

//         mid := (j+i)/2

// 		if isBadVersion(mid) {
// 			j = mid
// 		}else{
// 			i = mid
// 		}
// 	}

//     if isBadVersion(i)==true {
//         return i
//     }

// 	return j
// }
