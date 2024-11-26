package leetcode

func InorderTraverseBinaryTree(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	var result []int = make([]int, 0)
	result = append(result, InorderTraverseBinaryTree(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraverseBinaryTree(root.Right)...)
	return result
}

func InorderTraverseBinaryTreeNoRecursion(root *TreeNode) []int {

	return nil
}
