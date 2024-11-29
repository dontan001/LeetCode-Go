package leetcode

import "math"

func ValidateBinarySearchTree(root *TreeNode) bool {

	var isValid func(*TreeNode, int, int) bool
	isValid = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}

		v := node.Val
		if min < v && v < max && isValid(node.Left, min, v) && isValid(node.Right, v, max) {
			return true
		}

		return false
	}

	return isValid(root, math.MinInt64, math.MaxInt64)
}
