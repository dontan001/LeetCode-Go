package leetcode

func IsInterleave(s1 string, s2 string, s3 string) bool {

	var interleave func(int, int, int, bool) bool
	interleave = func(l int, r int, m int, left bool) bool {
		if l == len(s1) && r == len(s2) && m == len(s3) {
			return true
		}
		if l > len(s1) || r > len(s2) || m > len(s3) {
			return false
		}

		var step int
		for i := m; i < len(s3); i++ {
			if left {
				if l+step < len(s1) && s1[l+step] == s3[i] {
					step++
				} else {
					break
				}
			} else {
				if r+step < len(s2) && s2[r+step] == s3[i] {
					step++
				} else {
					break
				}
			}
		}

		for i := step; i > 0; i-- {
			if left {
				if interleave(l+i, r, m+i, false) {
					return true
				}
			} else {
				if interleave(l, r+i, m+i, true) {
					return true
				}
			}
		}
		return false
	}

	return interleave(0, 0, 0, true) || interleave(0, 0, 0, false)
}

func IsInterleaveDP(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
