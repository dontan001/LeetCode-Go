// Assume there is a function named generateTrees in the package
// which is the function we want to test.
package leetcode

import (
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	n := 3
	result := GenerateTrees(n)
	t.Logf("generateTrees(%d) = %#v", n, result)
}
