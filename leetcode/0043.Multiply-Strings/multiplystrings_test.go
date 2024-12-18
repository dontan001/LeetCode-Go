package leetcode

import "testing"

func TestMultiplyStrings(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		num1 := "2"
		num2 := "3"
		res := Multiply(num1, num2)
		t.Logf("multiply strings res: %v", res)
	})

	t.Run("case 2", func(t *testing.T) {
		num1 := "2000000000000000000000000000"
		num2 := "3000000000000000000000000000"
		res := Multiply(num1, num2)
		t.Logf("multiply strings res: %v", res)
	})
}
