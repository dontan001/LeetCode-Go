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
	top, botton := 0, row-1
	for top <= botton {
		m := (top + botton) / 2
		if m == top {
			if target >= matrix[botton][0] {
				targetRow = botton
			} else {
				targetRow = m
			}
			break
		}

		if target >= matrix[m][0] {
			top = m
		} else {
			botton = m
		}
	}

	if target > matrix[targetRow][col-1] {
		return false
	}
	left, right := 0, col-1
	for left <= right {
		m := (left + right) / 2
		if m == left {
			if target == matrix[targetRow][left] || target == matrix[targetRow][right] {
				return true
			}
			break
		}

		if target >= matrix[targetRow][m] {
			left = m
		} else {
			right = m
		}
	}

	return false
}
