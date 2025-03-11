package leetcode

func Partition(head *ListNode, x int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var hLeft, tLeft, hRight, tRight *ListNode
	var current *ListNode = head
	for current != nil {
		if current.Val < x {
			if hLeft == nil {
				hLeft = current
				tLeft = current
			} else {
				tLeft.Next = current
				tLeft = current
			}
		} else {
			if hRight == nil {
				hRight = current
				tRight = current
			} else {
				tRight.Next = current
				tRight = current
			}
		}
		current = current.Next
	}

	if tRight != nil {
		tRight.Next = nil
	}

	if tLeft != nil {
		tLeft.Next = hRight
		return hLeft
	} else {
		return hRight
	}
}
