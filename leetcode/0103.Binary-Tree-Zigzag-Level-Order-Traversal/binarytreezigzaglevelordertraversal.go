package leetcode

func ZigzagLevelOrder(root *TreeNode) [][]int {

	var results = [][]int{}
	if root == nil {
		return results
	}

	var queue, queueNext, lr = []*TreeNode{root}, []*TreeNode{}, true
	for len(queue) > 0 {
		var result = []int{}
		for i := len(queue) - 1; i >= 0; i-- {
			result = append(result, queue[i].Val)
			if lr {
				if queue[i].Left != nil {
					queueNext = append(queueNext, queue[i].Left)
				}
				if queue[i].Right != nil {
					queueNext = append(queueNext, queue[i].Right)
				}
			} else {
				if queue[i].Right != nil {
					queueNext = append(queueNext, queue[i].Right)
				}
				if queue[i].Left != nil {
					queueNext = append(queueNext, queue[i].Left)
				}
			}
		}
		results = append(results, result)
		queue, queueNext, lr = queueNext, []*TreeNode{}, !lr
	}

	return results
}
