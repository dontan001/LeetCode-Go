package leetcode

func RemoveNthNodeFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	var previous, current *ListNode = head, head
	for current.Next != nil {
		current = current.Next
		if n > 0 {
			n = n - 1
		} else {
			previous = previous.Next
		}
	}

	if n == 0 {
		if previous.Next != nil {
			previous.Next = previous.Next.Next
		}
	} else if n == 1 {
		head = head.Next
	}

	return head
}
