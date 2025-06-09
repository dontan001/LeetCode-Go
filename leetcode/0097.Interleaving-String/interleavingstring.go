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
