package leetcode

func BuildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	var rootIndex, rootIndexOther = 0, -1
	for i := 0; i < len(inorder); i++ {
		if preorder[rootIndex] == inorder[i] {
			rootIndexOther = i
			break
		}
	}

	var root = &TreeNode{Val: preorder[rootIndex]}
	root.Left = BuildTree(preorder[rootIndex+1:rootIndexOther+1], inorder[:rootIndexOther])
	root.Right = BuildTree(preorder[rootIndexOther+1:], inorder[rootIndexOther+1:])
	return root
}
