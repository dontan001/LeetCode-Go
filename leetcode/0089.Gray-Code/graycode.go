package leetcode

func GrayCode(n int) []int {

	var res []int = []int{}
	if n != 0 {
		res = append(res, 0)

		for i := 0; i < n; i++ {
			size := len(res)
			for j := size - 1; j >= 0; j-- {
				ele := res[j] | 1<<i
				res = append(res, ele)
			}
		}
	}
	return res
}
