package reverseKGroup

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummy
	end := dummy

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start, next := pre.Next, end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre
	}

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
