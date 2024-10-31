package leetcode

type Node struct {
	digit int
	next  *Node
}

func TwoAddV1(one *Node, another *Node) *Node {

	head := &Node{}
	currentHead := head

	currentOne := one
	currentAnother := another
	carry := 0
	for currentOne != nil || currentAnother != nil {
		if currentOne == nil {
			currentOne = &Node{0, nil}
		}
		if currentAnother == nil {
			currentAnother = &Node{0, nil}
		}

		currentDigit := currentOne.digit + currentAnother.digit + carry
		if currentDigit >= 10 {
			carry = 1
			currentDigit = currentDigit - 10
		} else {
			carry = 0
		}

		currentHead.digit = currentDigit
		currentHead.next = &Node{0, nil}
		if currentOne.next == nil && currentAnother.next == nil {
			currentHead.next.digit = carry
			break
		}

		currentHead = currentHead.next
		currentOne = currentOne.next
		currentAnother = currentAnother.next
	}

	return head
}
