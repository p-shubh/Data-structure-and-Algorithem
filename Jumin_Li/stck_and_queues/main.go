package main

import "fmt"

type Stack struct {
	Items []int
}

type Stack2 struct {
	Items []int
}

// Push
func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

// Pop will removce the value at the end
// and return the removed value
func (s *Stack) Pop() int {
	l := len(s.Items) - 1
	toRemove := s.Items[l]
	s.Items = s.Items[:l]
	return toRemove
}

// enquue
func (s *Stack2) Enquue(i int) {
	s.Items = append(s.Items, i)
}

// deque
func (s *Stack2) Deque() int {

	toRemove := s.Items[0]
	s.Items = s.Items[1:]

	return toRemove
}

func main() {
	myStack := Stack{}
	myStack2 := Stack2{}

	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	fmt.Println(myStack2)

	myStack2.Enquue(400)
	myStack2.Enquue(500)
	myStack2.Enquue(600)

	fmt.Println(myStack2)
	myStack2.Deque()
	fmt.Println(myStack2)

}
