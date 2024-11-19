package leetcode

func MergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {

	head := &ListNode{}

	var one, other, current *ListNode = list1, list2, head
	for one != nil && other != nil {
		if one.Val == other.Val {
			current.Next = &ListNode{Val: one.Val}
			current.Next.Next = &ListNode{Val: other.Val}
			one = one.Next
			other = other.Next
			current = current.Next.Next
		} else if one.Val < other.Val {
			current.Next = &ListNode{Val: one.Val}
			one = one.Next
			current = current.Next
		} else {
			current.Next = &ListNode{Val: other.Val}
			other = other.Next
			current = current.Next
		}
	}

	if one != nil {
		current.Next = one
	}
	if other != nil {
		current.Next = other
	}

	return head.Next
}
