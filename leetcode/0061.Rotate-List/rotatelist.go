package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RotateRight(head *ListNode, k int) *ListNode {

	var total int = 0
	current := head
	for current != nil {
		total++
		current = current.Next
	}

	if total == 0 {
		return head
	}
	k = k % total
	if k == 0 {
		return head
	}

	var i, j *ListNode = nil, head
	var distance int = 1
	for j.Next != nil {
		if distance < k {
			distance++
		} else {
			if i == nil {
				i = head
			} else {
				i = i.Next
			}
		}

		j = j.Next
	}

	if i != nil {
		tmp := i.Next
		i.Next = nil
		j.Next = head
		head = tmp
	}

	return head
}
