package leetcode

import "testing"

func TestGenerateParenthesesV1(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		res := GenerateParenthesesV1(4)
		t.Logf("there are %d combinations and they are \n %v", len(res), res)
	})
}
