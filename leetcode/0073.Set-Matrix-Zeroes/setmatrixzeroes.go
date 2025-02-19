package leetcode

func SetZeroes(matrix [][]int) {

	var row, col int = len(matrix), len(matrix[0])
	rowZeroes := make([]int, row)
	colZeroes := make([]int, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowZeroes[i] = 1
				colZeroes[j] = 1
			}
		}
	}

	for r, v := range rowZeroes {
		if v == 1 {
			for idx := 0; idx < col; idx++ {
				matrix[r][idx] = 0
			}
		}
	}

	for c, v := range colZeroes {
		if v == 1 {
			for idx := 0; idx < row; idx++ {
				matrix[idx][c] = 0
			}
		}
	}
}
