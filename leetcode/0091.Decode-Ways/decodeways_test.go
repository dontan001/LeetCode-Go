package leetcode

import "testing"

func TestNumDecodings(t *testing.T) {

	t.Run("test one", func(t *testing.T) {
		input := "12"
		want := 2
		actual := NumDecodings(input)
		if want != actual {
			t.Errorf("expected: %v got: %v", want, actual)
		}
	})
}
