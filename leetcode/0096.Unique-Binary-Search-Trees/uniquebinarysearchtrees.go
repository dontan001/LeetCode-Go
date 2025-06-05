package leetcode

func NumTrees(n int) int {
	if n <= 2 {
		return n
	}

	var countBSTrees func(int, int) int
	countBSTrees = func(start, end int) int {
		var temp int
		if start > end {
			return 1
		}

		for i := start; i <= end; i++ {
			left := countBSTrees(start, i-1)
			right := countBSTrees(i+1, end)
			temp += left * right
		}

		return temp
	}

	return countBSTrees(1, n)
}
