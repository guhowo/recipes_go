package linkedlist

func InitDoubleList() *Doublelist {
	list := Doublelist{Head: nil}
	return &list
}

func (l *Doublelist) Insert(value int) {
	node := &DoubleNode{value, nil, nil}
	node.Next = l.Head
	node.Pre = l.Head.Pre
	l.Head.Pre.Next = node
	l.Head.Pre = node
	l.Head = node
}
