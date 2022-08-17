package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {

	Second := l.Head
	l.Head = n
	l.Head.Next = Second
	l.Length++

}

func main() {

	mylist := LinkedList{}
	node1 := &Node{Data: 48}
	node2 := &Node{Data: 18}
	node3 := &Node{Data: 416}
	// node4 := &Node{Data: 48}

	mylist.Prepend(node1)
	mylist.Prepend(node2)
	mylist.Prepend(node3)

	fmt.Println(mylist)

	
}
