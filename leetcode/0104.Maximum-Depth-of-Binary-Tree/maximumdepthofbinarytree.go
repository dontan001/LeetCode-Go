package leetcode

func MaxDepth(root *TreeNode) int {

	var depth = 0
	if root == nil {
		return depth
	}

	var q, qNext = []*TreeNode{root}, []*TreeNode{}
	for len(q) != 0 {
		depth++
		for _, node := range q {
			if node.Left != nil {
				qNext = append(qNext, node.Left)
			}
			if node.Right != nil {
				qNext = append(qNext, node.Right)
			}
		}
		q, qNext = qNext, []*TreeNode{}
	}

	return depth
}
