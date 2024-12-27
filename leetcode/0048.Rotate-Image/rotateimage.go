package leetcode

func Rotate(matrix [][]int) {

	// matrix transpose
	row, column := len(matrix), len(matrix)
	for i := 0; i < row; i++ {
		for j := i + 1; j < column; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse
	for i := 0; i < row; i++ {
		begin, end := 0, column-1
		for begin < end {
			matrix[i][begin], matrix[i][end] = matrix[i][end], matrix[i][begin]
			begin++
			end--
		}
	}
}
