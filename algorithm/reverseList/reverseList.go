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

	return p1
}

func reverseListV2(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	var curNext *ListNode = nil

	for cur != nil {
		curNext = cur.Next
		cur.Next = pre
		pre = cur
		cur = curNext
	}

	return pre
}
