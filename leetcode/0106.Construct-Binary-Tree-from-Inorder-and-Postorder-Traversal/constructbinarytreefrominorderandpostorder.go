package leetcode

func BuildTree(inorder []int, postorder []int) *TreeNode {

	if len(inorder) == 0 {
		return nil
	}

	var rootIndex, rootIndexOther = len(postorder) - 1, -1
	root := &TreeNode{Val: postorder[rootIndex]}
	for index, node := range inorder {
		if node == root.Val {
			rootIndexOther = index
			break
		}
	}

	root.Left = BuildTree(inorder[:rootIndexOther], postorder[:rootIndexOther])
	root.Right = BuildTree(inorder[rootIndexOther+1:], postorder[rootIndexOther:rootIndex])
	return root
}
