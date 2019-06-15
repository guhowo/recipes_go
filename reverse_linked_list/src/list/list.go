package list

import (
	"fmt"
)

type Node struct {
	Element int
	Next *Node
}

type List struct {
	Head *Node
	Len int
}

var LinkedList = &List{}

func init() {
	LinkedList.Head = nil
	LinkedList.Len = 0
}

func (list *List)Add(e int) {
	node := &Node{
		e,
		list.Head,
	}
	list.Len++
	list.Head = node
}

func (list *List)Traversal() {
	ptr := list.Head
	for ptr != nil {
		fmt.Println(ptr.Element)
		ptr = ptr.Next
	}
}

func (list *List)Reverse() {
	if list.Len <= 1 {
		return
	}

	var head *Node
	ptr := list.Head
	ptr2 := ptr
	for ptr2 != nil {
		ptr2 = ptr.Next
		ptr.Next = head
		head = ptr
	}
}