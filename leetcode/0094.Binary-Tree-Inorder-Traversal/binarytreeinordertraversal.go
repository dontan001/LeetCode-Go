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

	var result []int = make([]int, 0)
	var stack []*TreeNode = make([]*TreeNode, 0)

	var cursor *TreeNode = root
	for cursor != nil || len(stack) != 0 {
		if cursor != nil {
			stack = append(stack, cursor)
			for cursor.Left != nil {
				stack = append(stack, cursor.Left)
				cursor = cursor.Left
			}
		}

		stackTop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, stackTop.Val)

		cursor = stackTop.Right
	}

	return result
}
