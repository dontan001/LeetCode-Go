package leetcode

func LengthOfLastWord(s string) int {

	var last, size int = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			size++
		} else {
			if size != 0 {
				last = size
			}
			size = 0
		}

		if i == len(s)-1 && size != 0 {
			last = size
		}
	}

	return last
}
