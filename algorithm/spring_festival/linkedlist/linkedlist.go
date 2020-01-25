//单链表
package linkedlist

func InitSinglelist() *SingleList {
	list := SingleList{Head: nil}
	return &list
}
func (list *SingleList) Insert(value int) {
	node := &Node{value, nil}
	node.Next = list.Head
	list.Head = node
}

func (list *SingleList) Delete(value int) {
	p1 := list.Head
	p2 := &Node{Next: p1}
	for p1 != nil {
		if p1.Elem == value {
			if p1 == list.Head {
				list.Head = list.Head.Next
			} else {
				p2.Next = p1.Next
			}
			break
		}
		p2 = p1
		p1 = p1.Next
	}
}
