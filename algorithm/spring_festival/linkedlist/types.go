package linkedlist

type List interface {
	Insert(int)
	Delete(int)
}

//单链表和循环链表的节点
type Node struct {
	Elem int
	Next *Node
}

//单链表
type SingleList struct {
	Head *Node
}

//循环链表
type CircleList struct {
	Head *Node
}

//双向链表的节点
type DoubleNode struct {
	Elem int
	Next *DoubleNode
	Pre  *DoubleNode
}

//双向链表
type Doublelist struct {
	Head *DoubleNode
}
