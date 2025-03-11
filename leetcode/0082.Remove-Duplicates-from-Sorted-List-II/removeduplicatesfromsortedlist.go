package leetcode

func DeleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var current, front, h, last *ListNode = head, head, nil, nil
	var counter int = 1
	for current != nil {
		current = current.Next
		if current != nil && current.Val == front.Val {
			counter++
		} else {
			if counter == 1 {
				if last == nil {
					last = front
					h = last
				} else {
					last.Next = front
					last = front
				}
			}

			front = current
			counter = 1
			if current == nil && last != nil {
				last.Next = nil
			}
		}
	}

	return h
}
