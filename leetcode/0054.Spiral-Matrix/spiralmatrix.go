package leetcode

func SpiralOrder(matrix [][]int) []int {

	result := make([]int, 0)

	rows, cols := len(matrix), len(matrix[0])
	x, y, dx, dy := 0, 0, 1, 0

	total := rows * cols
	for i := 0; i < total; i++ {
		result = append(result, matrix[y][x])
		matrix[y][x] = 101

		if !(0 <= x+dx && x+dx < cols && 0 <= y+dy && y+dy < rows) || matrix[y+dy][x+dx] == 101 {
			t := dx
			dx = -dy
			dy = t
		}

		x += dx
		y += dy
	}

	return result
}
