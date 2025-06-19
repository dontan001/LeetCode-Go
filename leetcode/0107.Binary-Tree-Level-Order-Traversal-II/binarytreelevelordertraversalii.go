package leetcode

func LevelOrderBottom(root *TreeNode) [][]int {

	var results = [][]int{}
	if root == nil {
		return results
	}

	var queue, queueNext = []*TreeNode{root}, []*TreeNode{}
	for len(queue) > 0 {
		var result = []int{}
		for _, node := range queue {
			result = append(result, node.Val)
			if node.Left != nil {
				queueNext = append(queueNext, node.Left)
			}
			if node.Right != nil {
				queueNext = append(queueNext, node.Right)
			}
		}

		var tmp = [][]int{result}
		results = append(tmp, results...)
		queue, queueNext = queueNext, []*TreeNode{}
	}

	return results
}
