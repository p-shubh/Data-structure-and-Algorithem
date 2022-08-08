package main

import (
	"fmt"
)

func main() {

	item := []int{1, 2, 3, 4, 5, 95, 78, 46, 58, 45, 86, 99, 251, 320}

	c, d := Linear_Search(item, 98)

	fmt.Println("print the aray ", c, d)

}

func Linear_Search(Data []int, value int) (bool, int) {

	var data_list int

	var check bool = false

	for _, i := range Data {

		if i == value {
			check = true
			data_list = value

		}
	}
	return check, data_list
}
