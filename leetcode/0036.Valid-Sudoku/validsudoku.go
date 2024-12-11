package leetcode

import "strconv"

func ValidSudoku(board [][]byte) bool {

	appears := make(map[string]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			number := string(board[i][j])
			if number != "." {
				row := number + "row" + strconv.Itoa(i)
				col := number + "col" + strconv.Itoa(j)
				grid := number + "grid" + strconv.Itoa(i/3) + strconv.Itoa(j/3)
				if appears[row] || appears[col] || appears[grid] {
					return false
				}

				appears[row] = true
				appears[col] = true
				appears[grid] = true
			}
		}
	}

	return true
}
