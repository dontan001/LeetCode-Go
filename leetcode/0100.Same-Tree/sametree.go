package leetcode

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !IsSameTree(p.Left, q.Left) {
		return false
	}
	if !IsSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
