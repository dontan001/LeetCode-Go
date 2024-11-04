package leetcode

func TwoAddV1(one *ListNode, another *ListNode) *ListNode {

	head := &ListNode{}
	currentHead := head

	currentOne := one
	currentAnother := another
	carry := 0
	for currentOne != nil || currentAnother != nil {
		if currentOne == nil {
			currentOne = &ListNode{0, nil}
		}
		if currentAnother == nil {
			currentAnother = &ListNode{0, nil}
		}

		currentDigit := currentOne.Val + currentAnother.Val + carry
		if currentDigit >= 10 {
			carry = 1
			currentDigit = currentDigit - 10
		} else {
			carry = 0
		}

		currentHead.Val = currentDigit
		currentHead.Next = &ListNode{0, nil}
		if currentOne.Next == nil && currentAnother.Next == nil {
			currentHead.Next.Val = carry
			break
		}

		currentHead = currentHead.Next
		currentOne = currentOne.Next
		currentAnother = currentAnother.Next
	}

	return head
}
