package leetcode

import "testing"

func TestGenerateMatrix(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		matrix := GenerateMatrix(3)
		t.Logf("matrix = %v", matrix)
	})

	t.Run("case 2", func(t *testing.T) {
		matrix := GenerateMatrix(4)
		t.Logf("matrix = %v", matrix)
	})
}
