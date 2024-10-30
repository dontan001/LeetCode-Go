package main

func include(source string, sub string) bool {
	var i, j int
	for i = 0; i < len(source); i++ {
		if source[i] == sub[0] {
			for j = 0; j < len(sub); j++ {
				if source[i+j] != sub[j] {
					break
				}

				if j == len(sub)-1 {
					return true
				}
			}
		}

		if len(source)-(i+1) < len(sub) {
			break
		}
	}

	return false
}
