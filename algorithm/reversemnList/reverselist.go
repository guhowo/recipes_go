package reversemnList

//给定一个连表，从m到n反转

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode, m, n int) {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummy
	cur := dummy.Next

	for {
		n := k
	}

}
