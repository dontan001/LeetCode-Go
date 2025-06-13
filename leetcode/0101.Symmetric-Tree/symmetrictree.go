package leetcode

func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var mirror func(*TreeNode, *TreeNode) bool
	mirror = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		return mirror(left.Left, right.Right) && mirror(left.Right, right.Left)
	}

	return mirror(root.Left, root.Right)
}
