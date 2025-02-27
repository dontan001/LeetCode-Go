package leetcode

func Exist(board [][]byte, word string) bool {

	var left, right, up, down int = -1, 1, 2, -2
	var row, col int = len(board), len(board[0])

	var visit [][]int = make([][]int, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]int, col)
	}

	var exist bool
	var search func(int, int, int, int)
	search = func(start int, r int, c int, blocked int) {
		if start == len(word) {
			exist = true
			return
		}

		if !(r >= 0 && r < row && c >= 0 && c < col) {
			return
		}

		if visit[r][c] == 1 {
			return
		}

		if word[start] != board[r][c] {
			return
		}

		visit[r][c] = 1
		if !exist && (blocked == 0 || blocked != right) {
			search(start+1, r, c+1, left)
		}

		if !exist && (blocked == 0 || blocked != left) {
			search(start+1, r, c-1, right)
		}

		if !exist && (blocked == 0 || blocked != down) {
			search(start+1, r-1, c, up)
		}

		if !exist && (blocked == 0 || blocked != up) {
			search(start+1, r+1, c, down)
		}

		if !exist {
			visit[r][c] = 0
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			search(0, i, j, 0)
			if exist {
				return true
			}
		}
	}

	return exist
}
