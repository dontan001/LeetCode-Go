package leetcode

func GenerateParenthesesV1(n int) []string {

	result := []string{}

	var generateParentheses func(int, int, string)
	generateParentheses = func(leftCount, rightCount int, prefix string) {
		if leftCount == 0 && rightCount == 0 {
			result = append(result, prefix)
			return
		}

		if leftCount > 0 {
			generateParentheses(leftCount-1, rightCount, prefix+"(")
		}

		if rightCount > 0 && leftCount < rightCount {
			generateParentheses(leftCount, rightCount-1, prefix+")")
		}
	}

	generateParentheses(n, n, "")
	return result
}

// wrong answer
func GenerateParenthesesV2(n int) []string {

	if n == 1 {
		return []string{"()"}
	}

	parentheses := GenerateParenthesesV2(n - 1)
	newParentheses := []string{}
	exists := make(map[string]bool)
	for _, parenthese := range parentheses {
		left := "()" + parenthese
		if !exists[left] {
			newParentheses = append(newParentheses, left)
			exists[left] = true
		}

		right := parenthese + "()"
		if !exists[right] {
			newParentheses = append(newParentheses, right)
			exists[right] = true
		}

		surround := "(" + parenthese + ")"
		if !exists[surround] {
			newParentheses = append(newParentheses, surround)
			exists[surround] = true
		}
	}

	return newParentheses
}
