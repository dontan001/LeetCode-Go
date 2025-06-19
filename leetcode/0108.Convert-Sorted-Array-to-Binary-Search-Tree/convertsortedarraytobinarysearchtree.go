package leetcode

func SortedArrayToBST(nums []int) *TreeNode {

	var buildtree func(int, int) *TreeNode
	buildtree = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}
		if start == end {
			return &TreeNode{Val: nums[start]}
		}

		var middle = (start + end) / 2
		var root = &TreeNode{Val: nums[middle]}
		root.Left = buildtree(start, middle-1)
		root.Right = buildtree(middle+1, end)
		return root
	}

	return buildtree(0, len(nums)-1)
}
