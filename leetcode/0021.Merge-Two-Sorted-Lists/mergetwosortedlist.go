package leetcode

func MergeTwoSortedListV3(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = MergeTwoSortedListV3(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoSortedListV3(list2.Next, list1)
		return list2
	}
}

func MergeTwoSortedListV2(list1 *ListNode, list2 *ListNode) *ListNode {

	head := &ListNode{}
	current := head
	mergeList(list1, list2, current)
	return head.Next
}

func mergeList(list1 *ListNode, list2 *ListNode, current *ListNode) {
	if list1 == nil {
		current.Next = list2
		return
	}

	if list2 == nil {
		current.Next = list1
		return
	}

	if list1.Val == list2.Val {
		current.Next = &ListNode{Val: list1.Val}
		current.Next.Next = &ListNode{Val: list2.Val}
		mergeList(list1.Next, list2.Next, current.Next.Next)
	} else if list1.Val > list2.Val {
		current.Next = &ListNode{Val: list2.Val}
		mergeList(list1, list2.Next, current.Next)
	} else {
		current.Next = &ListNode{Val: list1.Val}
		mergeList(list1.Next, list2, current.Next)
	}
}

func MergeTwoSortedListsV1(list1 *ListNode, list2 *ListNode) *ListNode {

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
