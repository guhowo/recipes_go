package linkedlist

//带头节点的list

func InitCircleList() *CircleList {
	list := CircleList{Head: &Node{}}
	list.Head.Next = list.Head
	return &list
}

func (list *CircleList) Insert(value int) {
	node := &Node{value, nil}
	node.Next = list.Head.Next
	list.Head.Next = node
}

func (list *CircleList) Delete(value int) {
	pre := list.Head.Next
	p := list.Head
	for pre != list.Head {
		if pre.Elem == value {
			p.Next = pre.Next
			return
		}
		p = pre
		pre = pre.Next
	}
}
