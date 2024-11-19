package leetcode

func ValidateParentheses(s string) bool {
	match := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		}

		if char == ')' || char == ']' || char == '}' {
			if len(stack) <= 0 || match[stack[len(stack)-1]] != char {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
