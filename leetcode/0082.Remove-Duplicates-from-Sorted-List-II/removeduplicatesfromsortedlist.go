package leetcode

func DeleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var current, previous, h, tail *ListNode = head, head, nil, nil
	var counter int = 1
	for current != nil {
		current = current.Next
		if current != nil && current.Val == previous.Val {
			counter++
		} else {
			if counter == 1 {
				if tail == nil {
					tail = previous
					h = tail
				} else {
					tail.Next = previous
					tail = previous
				}
			}

			previous = current
			counter = 1
			if current == nil && tail != nil {
				tail.Next = nil
			}
		}
	}

	return h
}
