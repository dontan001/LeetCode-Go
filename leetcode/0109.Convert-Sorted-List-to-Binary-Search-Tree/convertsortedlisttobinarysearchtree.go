package leetcode

func SortedListToBST(head *ListNode) *TreeNode {

	var buildtree func(*ListNode, *ListNode) *TreeNode
	buildtree = func(begin, end *ListNode) *TreeNode {

		if begin == end {
			return nil
		}

		var slow, fast *ListNode = begin, begin
		for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
			if fast.Next == end || fast.Next.Next == end {
				break
			}
			slow = slow.Next
			fast = fast.Next.Next
		}

		root := &TreeNode{Val: slow.Val}
		root.Left = buildtree(begin, slow)
		root.Right = buildtree(slow.Next, end)
		return root
	}

	return buildtree(head, nil)
}
