package leetcode

func SwapNodesInPairs(head *ListNode) *ListNode {

	var p, c *ListNode = nil, head
	var counter int
	for c != nil {
		counter = counter + 1
		if counter%2 == 0 {
			c.Val, p.Val = p.Val, c.Val
		}

		p = c
		c = c.Next
	}

	return head
}
