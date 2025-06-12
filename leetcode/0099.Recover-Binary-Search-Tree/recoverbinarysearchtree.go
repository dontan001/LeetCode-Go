package leetcode

/*
Three cases for inorder traversal :
case 1 : All are in ascending order
No need to do anything
example - [1, 2, 3, 4, 5]

case 2 : one pair is in descending order
swap both pairs
example - [2, 1, 3, 4, 5]

case 3 : two pairs are in descending order
swap first pair's first value and second pair's second value
example - [5, 2, 3, 4, 1]
*/
func RecoverTree(root *TreeNode) {

	var prev, first, second *TreeNode
	var inorder func(*TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}

		inorder(r.Left)
		if prev != nil && prev.Val > r.Val {
			if first == nil {
				first = prev
			}
			second = r
		}
		prev = r
		inorder(r.Right)
	}

	inorder(root)
	first.Val, second.Val = second.Val, first.Val
}
