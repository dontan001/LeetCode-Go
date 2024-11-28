package leetcode

import "testing"

func TestInorderTraverseBinaryTree(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		root3 := &TreeNode{Val: 3, Left: nil, Right: nil}
		root2 := &TreeNode{Val: 2, Left: root3, Right: nil}
		root := &TreeNode{Val: 1, Left: nil, Right: root2}

		result := InorderTraverseBinaryTreeNoRecursion(root)
		t.Logf("result is : %v", result)
	})
}
