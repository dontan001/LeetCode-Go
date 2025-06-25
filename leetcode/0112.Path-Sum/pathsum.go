package leetcode

func HasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return HasPathSum(root.Left, targetSum-root.Val) ||
		HasPathSum(root.Right, targetSum-root.Val)
}
