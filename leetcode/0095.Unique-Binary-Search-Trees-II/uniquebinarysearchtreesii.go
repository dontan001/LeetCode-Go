package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func GenerateTrees(n int) []*TreeNode {

	var generate func(int, int) []*TreeNode
	generate = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}

		trees := []*TreeNode{}
		for i := start; i <= end; i++ {
			left := generate(start, i-1)
			right := generate(i+1, end)
			for _, l := range left {
				for _, r := range right {
					root := &TreeNode{Val: i, Left: l, Right: r}
					trees = append(trees, root)
				}
			}
		}

		return trees
	}

	return generate(1, n)
}
