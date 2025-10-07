package leetcode

func Flatten(root *TreeNode) {

	var nodes []*TreeNode = []*TreeNode{}
	var preoder func(root *TreeNode)
	preoder = func(root *TreeNode) {
		if root == nil {
			return
		}
		nodes = append(nodes, root)
		preoder(root.Left)
		preoder(root.Right)
	}

	preoder(root)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
}
