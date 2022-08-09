package main

import (
	"fmt"
	"sort"
)

func main() {

	item := []int{1, 2, 3, 4, 5, 95, 78, 46, 58, 45, 86, 99, 251, 320}

	fmt.Println("item", item)
	sort.Sort(sort.Reverse(sort.IntSlice(item)))

	fmt.Println("item", item)
	sort.Ints(item)

	fmt.Println("item", item)

	c, d := Binary_Search(item, item[6])

	fmt.Println("print binary aray ", c, d)
	// fmt.Println("print binary aray ", 7/2)
}

func Binary_Search(data_list []int, needle int) (bool, int) {

	result := true

	low := 0

	high := len(data_list) - 1

	for low <= high {
		median := (low + high) / 2

		if data_list[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(data_list) || data_list[low] != needle {
		result = false
		return result, needle
	}
	return result, needle
}
