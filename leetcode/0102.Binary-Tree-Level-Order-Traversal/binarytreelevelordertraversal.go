package leetcode

func LevelOrder(root *TreeNode) [][]int {

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
		results = append(results, result)
		queue, queueNext = queueNext, []*TreeNode{}
	}
	return results
}
