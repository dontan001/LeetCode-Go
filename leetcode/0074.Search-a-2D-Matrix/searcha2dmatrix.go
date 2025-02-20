package leetcode

func SearchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[row-1][col-1] {
		return false
	}

	var targetRow int = -1
	begin, end := 0, row-1
	for begin <= end {
		m := (begin + end) / 2
		if m == begin {
			if target >= matrix[end][0] {
				targetRow = end
			} else {
				targetRow = m
			}
			break
		}

		if target >= matrix[m][0] {
			begin = m
		} else {
			end = m
		}
	}

	if target > matrix[targetRow][col-1] {
		return false
	}
	begin, end = 0, col-1
	for begin <= end {
		m := (begin + end) / 2
		if m == begin {
			if target == matrix[targetRow][begin] || target == matrix[targetRow][end] {
				return true
			}
			break
		}

		if target >= matrix[targetRow][m] {
			begin = m
		} else {
			end = m
		}
	}

	return false
}
