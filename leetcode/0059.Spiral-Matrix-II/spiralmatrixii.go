package leetcode

func GenerateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var x, y, dx, dy = 0, 0, 1, 0
	for i := 0; i < n*n; i++ {
		matrix[y][x] = i + 1
		if !(0 <= x+dx && x+dx < n && 0 <= y+dy && y+dy < n) || matrix[y+dy][x+dx] != 0 {
			dx, dy = -dy, dx
		}

		x += dx
		y += dy
	}

	return matrix
}
