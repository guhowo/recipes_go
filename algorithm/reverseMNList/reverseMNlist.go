package reverseMNList

//给定一个连表，从m到n反转

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummy
	end := dummy

	for i := 1; i < m && pre.Next != nil; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	start := pre.Next
	end = start

	for i := 0; i < n-m; i++ {
		end = end.Next
	}

	next := end.Next
	end.Next = nil
	pre.Next = reverse(start)
	start.Next = next

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre, curNext *ListNode = nil, nil

	for cur != nil {
		curNext = cur.Next
		cur.Next = pre
		pre = cur
		cur = curNext
	}

	return pre
}
