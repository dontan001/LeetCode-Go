package leetcode

import "testing"

func TestReverseInteger(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		a := -10
		t.Logf("divide: %d , %d", a/10, a%10)
	})

	t.Run("case 2", func(t *testing.T) {
		a := 120
		after := ReverseInteger(a)
		t.Logf("divide: %d", after)
	})

	t.Run("case 3", func(t *testing.T) {
		a := -120
		after := ReverseInteger(a)
		t.Logf("divide: %d", after)
	})
}
