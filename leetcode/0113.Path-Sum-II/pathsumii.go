package leetcode

func PathSum(root *TreeNode, targetSum int) [][]int {

	var paths = [][]int{}
	var backtrack func(*TreeNode, int, []int)
	backtrack = func(root *TreeNode, targetSum int, path []int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if targetSum == root.Val {
				path = append(path, root.Val)
				paths = append(paths, append([]int{}, path...))
			}
			return
		}

		path = append(path, root.Val)
		backtrack(root.Left, targetSum-root.Val, path)
		backtrack(root.Right, targetSum-root.Val, path)
		path = path[:len(path)-1]
	}

	backtrack(root, targetSum, []int{})
	return paths
}
