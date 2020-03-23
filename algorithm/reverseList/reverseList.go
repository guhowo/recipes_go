package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

//单链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	p1.Next = nil
	for p2 != nil {
		pre := p1
		p1 = p2
		p2 = p2.Next
		p1.Next = pre
	}

	head = p1
	return p1
}
