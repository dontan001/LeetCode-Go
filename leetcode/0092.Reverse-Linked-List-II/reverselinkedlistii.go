package leetcode

func ReverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil || left == right {
		return head
	}

	var dummy = &ListNode{Next: head}
	var p = dummy
	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	var c = p.Next
	for i := 0; i < right-left; i++ {
		var t = c.Next
		c.Next = t.Next
		t.Next = p.Next
		p.Next = t
	}

	return dummy.Next
}
