package leetcode

import "testing"

func TestValidateParentheses(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		s := "()"
		res := ValidateParentheses(s)
		t.Log(res)
	})
}
